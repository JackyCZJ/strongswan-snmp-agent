# strongswan-snmp-agent

Play For fun.

Use vici api transfer strongSwan data  to SNMP.

use cisco's ipsec mib , but not complete.

Use own mib if you want to avoid the cisco's mib.

known issue:
- strongswan data not fully compatible with cisco's mib.
- a lot of todo.
- not test in production.
- snmpwalk will report error for not sorted oid. use -Cc to ignore it.
    
    