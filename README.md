# Permission
Define some standard permissions
- ActionNone     int32 = 0
- ActionView     int32 = 1
- ActionSearch   int32 = 2
- ActionQueryAll int32 = 4
- ActionViewAll  int32 = 256 - 1
- ActionAdd      int32 = 256
- ActionEdit     int32 = 512
- ActionPatch    int32 = 1024
- ActionApprove  int32 = 2048
- ActionReject   int32 = 4096
- ActionDelete   int32 = 32768
- ActionWriteAll int32 = 65536 - 1
- ActionAll      int32 = 2147483647

## Installation

Please make sure to initialize a Go module before installing common-go/permission:

```shell
go get -u github.com/common-go/permission
```

Import:

```go
import "github.com/common-go/permission"
```

#### You can optimize the import by version:
- v0.0.4: Read Write Approve Delete
- v0.0.5: Read Write Delete Approve
- v0.0.6: Read Add Edit Approve Delete
- v0.0.7: Read Add Edit Delete Approve
- v0.0.8: Read Add Edit Patch Approve Delete
- v0.0.9: Read Add Edit Patch Delete Approve