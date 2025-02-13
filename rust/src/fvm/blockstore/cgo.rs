use std::ptr;

use anyhow::{anyhow, Result};
use cid::Cid;
use fvm_ipld_blockstore::Blockstore;

use super::super::cgo::*;

/// The maximum amount of data to buffer in a batch before writing it to the underlying blockstore.
const MAX_BUF_SIZE: usize = 4 << 20; // 4MiB
/// The maximum number of blocks to buffer in a batch before before writing it to the underlying
/// blockstore.
const MAX_BLOCK_BATCH: usize = 1024;

/// A rough estimate of the CID size, used to estimate the maximum amount of space much space we'll
/// need in the batch buffer to store a CID.
const EST_MAX_CID_LEN: usize = 100;

pub struct CgoBlockstore {
    handle: u64,
}

impl CgoBlockstore {
    /// Construct a new blockstore from a handle.
    pub fn new(handle: u64) -> CgoBlockstore {
        CgoBlockstore { handle }
    }
}

impl Blockstore for CgoBlockstore {
    fn has(&self, k: &Cid) -> Result<bool> {
        let k_bytes = k.to_bytes();
        unsafe {
            match cgo_blockstore_has(self.handle, k_bytes.as_ptr(), k_bytes.len() as i32) {
                // We shouldn't get an "error not found" here, but there's no reason to be strict
                // about it.
                0 | ERR_NOT_FOUND => Ok(false),
                1 => Ok(true),
                // Panic on unknown values. There's a bug in the program.
                r @ 2.. => panic!("invalid return value from has: {}", r),
                // Panic if the store isn't registered. This means something _very_ unsafe is going
                // on and there is a bug in the program.
                ERR_INVALID_HANDLE => panic!("blockstore {} not registered", self.handle),
                // Otherwise, return "other". We should add error codes in the future.
                e => Err(anyhow!("cgo blockstore 'has' failed with error code {}", e)),
            }
        }
    }

    fn get(&self, k: &Cid) -> Result<Option<Vec<u8>>> {
        let k_bytes = k.to_bytes();
        unsafe {
            let mut buf: *mut u8 = ptr::null_mut();
            let mut size: i32 = 0;
            match cgo_blockstore_get(
                self.handle,
                k_bytes.as_ptr(),
                k_bytes.len() as i32,
                &mut buf,
                &mut size,
            ) {
                0 => Ok(Some(Vec::from_raw_parts(buf, size as usize, size as usize))),
                r @ 1.. => panic!("invalid return value from get: {}", r),
                ERR_INVALID_HANDLE => panic!("blockstore {} not registered", self.handle),
                ERR_NOT_FOUND => Ok(None),
                e => Err(anyhow!("cgo blockstore 'get' failed with error code {}", e)),
            }
        }
    }

    fn put_many_keyed<D, I>(&self, blocks: I) -> Result<()>
    where
        Self: Sized,
        D: AsRef<[u8]>,
        I: IntoIterator<Item = (Cid, D)>,
    {
        fn flush_buffered(handle: u64, lengths: &mut Vec<i32>, buf: &mut Vec<u8>) -> Result<()> {
            if buf.is_empty() {
                return Ok(());
            }

            unsafe {
                let result = cgo_blockstore_put_many(
                    handle,
                    lengths.as_ptr(),
                    lengths.len() as i32,
                    buf.as_ptr(),
                );
                buf.clear();
                lengths.clear();

                match result {
                    0 => Ok(()),
                    r @ 1.. => panic!("invalid return value from put_many: {}", r),
                    ERR_INVALID_HANDLE => panic!("blockstore {} not registered", handle),
                    // This error makes no sense.
                    ERR_NOT_FOUND => panic!("not found error on put"),
                    e => Err(anyhow!("cgo blockstore 'put' failed with error code {}", e)),
                }
            }
        }

        let mut lengths = Vec::with_capacity(MAX_BLOCK_BATCH);
        let mut buf = Vec::with_capacity(MAX_BUF_SIZE);
        for (k, block) in blocks {
            let block = block.as_ref();
            // We limit both the max number of blocks and the max buffer size. Technically, we could
            // _just_ limit the buffer size as that should bound the number of blocks. However,
            // bounding the maximum number of blocks means we can allocate the vector up-front and
            // avoids any re-allocation, copying, etc.
            if lengths.len() >= MAX_BLOCK_BATCH
                || EST_MAX_CID_LEN + block.len() + buf.len() > MAX_BUF_SIZE
            {
                flush_buffered(self.handle, &mut lengths, &mut buf)?;
            }

            let start = buf.len();
            k.write_bytes(&mut buf)?;
            buf.extend_from_slice(block);
            let size = buf.len() - start;
            lengths.push(size as i32);
        }
        flush_buffered(self.handle, &mut lengths, &mut buf)
    }

    fn put_keyed(&self, k: &Cid, block: &[u8]) -> Result<()> {
        let k_bytes = k.to_bytes();
        unsafe {
            match cgo_blockstore_put(
                self.handle,
                k_bytes.as_ptr(),
                k_bytes.len() as i32,
                block.as_ptr(),
                block.len() as i32,
            ) {
                0 => Ok(()),
                r @ 1.. => panic!("invalid return value from put: {}", r),
                ERR_INVALID_HANDLE => panic!("blockstore {} not registered", self.handle),
                // This error makes no sense.
                ERR_NOT_FOUND => panic!("not found error on put"),
                e => Err(anyhow!("cgo blockstore 'put' failed with error code {}", e)),
            }
        }
    }
}
