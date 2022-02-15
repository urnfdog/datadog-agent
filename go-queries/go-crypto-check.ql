/**
 * @name Go Crypto Check
 * @description Reveals usages of go/crypto for FIPS compliance purposes
 * @kind problem
 * @precision very-high
 * @id go/crypto-fips-check
 * @tags security
 * @security-severity 10.0
 */

import go

from ImportSpec i
where i.getPath().regexpMatch("crypto/.*")
select i, "Non-FIPS compliant go crypto import was found: " + i.getPath()
