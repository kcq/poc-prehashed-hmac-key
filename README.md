# PoC - Pre-hashed HMAC Key

When your HMAC key is bigger than your hashing function block size you can pre-hash the key. This HMAC behavior is useful because now you no longer have to store the key itself.

## Output

SHA256 Block Size: 64
key -> secret.key:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA (65)
keyPH -> a14ab6d394f8e48b497a207c4b70487e8de4c5ad42c56999f16ef0303deccb680000000000000000000000000000000000000000000000000000000000000000 (64)
HMAC(key,data)   -> fede978063a6d53f7d3c01710ff67814a29a8496ca1a55eec7b9c5ceb762a4be
HMAC(keyPH,data) -> fede978063a6d53f7d3c01710ff67814a29a8496ca1a55eec7b9c5ceb762a4be
HMAC Match!




