//! This module contains the "mock" version of the cgo externs. We need these so that cbindgen
//! doesn't try to put them in the C header. Otherwise, c-for-go can't parse the header.

pub fn cgo_blockstore_get(
    store: u64,
    k: *const u8,
    k_len: i32,
    block: *mut *mut u8,
    size: *mut i32,
) -> i32 {
    unimplemented!()
}

pub fn cgo_blockstore_put(
    store: u64,
    k: *const u8,
    k_len: i32,
    block: *const u8,
    block_len: i32,
) -> i32 {
    unimplemented!()
}

pub fn cgo_blockstore_put_many(
    store: u64,
    lengths: *const i32,
    lengths_len: i32,
    blocks: *const u8,
) -> i32 {
    unimplemented!()
}

pub fn cgo_blockstore_has(store: u64, k: *const u8, k_len: i32) -> i32 {
    unimplemented!()
}

pub fn cgo_extern_get_chain_randomness(
    handle: u64,
    pers: i64,
    round: i64,
    entropy: *const u8,
    entropy_len: i32,
    randomness: *mut [u8; 32],
) -> i32 {
    unimplemented!()
}

pub fn cgo_extern_get_beacon_randomness(
    handle: u64,
    pers: i64,
    round: i64,
    entropy: *const u8,
    entropy_len: i32,
    randomness: *mut [u8; 32],
) -> i32 {
    unimplemented!()
}

pub fn cgo_extern_verify_consensus_fault(
    handle: u64,
    h1: *const u8,
    h1_len: i32,
    h2: *const u8,
    h2_len: i32,
    extra: *const u8,
    extra_len: i32,
    miner_id: *mut u64,
    epoch: *mut i64,
    fault: *mut i64,
    gas_used: *mut i64,
) -> i32 {
    unimplemented!()
}
