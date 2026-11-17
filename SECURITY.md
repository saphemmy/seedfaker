# Security

## Not a cryptographic generator

seedfaker is deterministic and predictable by design. Its output must never be used
for anything that needs to be unguessable — passwords, tokens, session IDs, keys, or
secrets of any kind. Use `crypto/rand` for those.

Generated emails use reserved domains (RFC 2606) and can never reach a real inbox.
Other generated data is synthetic, but it is not guaranteed to avoid matching a real
person, company, address or identifier by coincidence. Do not treat it as anonymised
production data.

## Reporting a vulnerability

If you find a security issue in the library itself, open a report through the
repository's private vulnerability reporting, or a regular issue if it is not
sensitive. Please do not include real secrets in any report.
