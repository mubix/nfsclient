# nfsclient

NFS client written in Go, and has ls, upload, download, rm, mkdir, and rmdir capabilities

```
nfsclient.exe <host>:<target path> <access level root:0:0> <command ls/up/down/rm/mkdir/rmdir> <path if required> <dest if upload>
```

### ls

```
PS C:\temp> .\nfsclient.exe 192.168.1.100:/export/share root:0:0 ls ./
2022/10/29 01:47:35 5
[C:\temp\nfsclient.exe 192.168.1.100:/export/share root:0:0 ls ./]
2022/10/29 01:47:35 host=192.168.1.100 target=/export/share command=ls
2022/10/29 01:47:35 Connecting to 192.168.1.100:892 from unprivileged port
2022/10/29 01:47:35 Connecting to 192.168.1.100:2049 from unprivileged port
2022/10/29 01:47:35 192.168.1.100:/export/share fsinfo=&nfs.FSInfo{Attr:nfs.PostOpAttr{IsSet:false, Attr:nfs.Fattr{Type:0x0, FileMode:0x0, Nlink:0x0, UID:0x0, GID:0x0, Filesize:0x0, Used:0x0, SpecData:[2]uint32{0x0, 0x0}, FSID:0x0, Fileid:0x0, Atime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Mtime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Ctime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}}}, RTMax:0x20000, RTPref:0x20000, RTMult:0x1000, WTMax:0x20000, WTPref:0x20000, WTMult:0x1000, DTPref:0x1000, Size:0xffffffff000, TimeDelta:nfs.NFS3Time{Seconds:0x1, Nseconds:0x0}, Properties:0x1b}
2022/10/29 01:47:35 root -> 0x0100070080e3a60100000000fad64994ded0c3610000000000000000
2022/10/29 01:47:35 No EOF for dirents so calling back for more
2022/10/29 01:47:35 dirs:
+--------------------------+------+-----+----------+-----------+
| FILENAME                 |  UID | GID | MODE     |      SIZE |
+--------------------------+------+-----+----------+-----------+
| ..                       |    0 |   0 | 0x588240 |      4096 |
| .                        |    0 |   0 | 0x588240 |      4096 |
| DB                       | 1024 | 100 | 0x588240 |      4096 |
| OUT                      | 1024 | 100 | 0x588240 |      4096 |
| nfs.go                   | 1024 | 100 | 0x588240 |        36 |
| Test1.txt                | 1024 | 100 | 0x588240 |      8196 |
+--------------------------+------+-----+----------+-----------+
2022/10/29 01:47:35 Completed tests
```

### upload

```
PS C:\temp> .\nfsclient.exe 192.168.1.100:/export/share root:0:0 up .\output.txt ./redteam/output.txt
2022/10/29 01:54:36 6
[C:\temp\nfsclient.exe 192.168.1.100:/export/share root:0:0 up .\output.txt ./redteam/output.txt]
2022/10/29 01:54:36 host=192.168.1.100 target=/export/share command=up
2022/10/29 01:54:36 Connecting to 192.168.1.100:892 from unprivileged port
2022/10/29 01:54:36 Connecting to 192.168.1.100:2049 from unprivileged port
2022/10/29 01:54:36 192.168.1.100:/export/share fsinfo=&nfs.FSInfo{Attr:nfs.PostOpAttr{IsSet:false, Attr:nfs.Fattr{Type:0x0, FileMode:0x0, Nlink:0x0, UID:0x0, GID:0x0, Filesize:0x0, Used:0x0, SpecData:[2]uint32{0x0, 0x0}, FSID:0x0, Fileid:0x0, Atime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Mtime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Ctime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}}}, RTMax:0x20000, RTPref:0x20000, RTMult:0x1000, WTMax:0x20000, WTPref:0x20000, WTMult:0x1000, DTPref:0x1000, Size:0xffffffff000, TimeDelta:nfs.NFS3Time{Seconds:0x1, Nseconds:0x0}, Properties:0x1b}
2022/10/29 01:54:36 Opening .\output.txt
2022/10/29 01:54:36 lookup(redteam): FH 0x0100070180e3a60100000000fad64994ded0c36100000000000000000380a6014adc6b01, attr: {Type:2 FileMode:511 Nlink:2 UID:0 GID:0 Filesize:4096 Used:4096 SpecData:[0 0] FSID:7044703896526771962 Fileid:27688963 Atime:{Seconds:1667022834 Nseconds:221578598} Mtime:{Seconds:1667022834 Nseconds:221578598} Ctime:{Seconds:1667022834 Nseconds:238578915}}
2022/10/29 01:54:36 lookup(output.txt): file does not exist
2022/10/29 01:54:36 lookup(redteam): FH 0x0100070180e3a60100000000fad64994ded0c36100000000000000000380a6014adc6b01, attr: {Type:2 FileMode:511 Nlink:2 UID:0 GID:0 Filesize:4096 Used:4096 SpecData:[0 0] FSID:7044703896526771962 Fileid:27688963 Atime:{Seconds:1667022834 Nseconds:221578598} Mtime:{Seconds:1667022834 Nseconds:221578598} Ctime:{Seconds:1667022834 Nseconds:238578915}}
2022/10/29 01:54:36 create(./redteam/output.txt): created successfully
2022/10/29 01:54:36 write(0100070280e3a60100000000fad64994ded0c36100000000000000000580a6014bdc6b010380a6014adc6b01) len=1916 new_offset=1916 written=1916 total=1916
2022/10/29 01:54:36 lookup(redteam): FH 0x0100070180e3a60100000000fad64994ded0c36100000000000000000380a6014adc6b01, attr: {Type:2 FileMode:511 Nlink:2 UID:0 GID:0 Filesize:4096 Used:4096 SpecData:[0 0] FSID:7044703896526771962 Fileid:27688963 Atime:{Seconds:1667022834 Nseconds:221578598} Mtime:{Seconds:1667022876 Nseconds:651371503} Ctime:{Seconds:1667022876 Nseconds:651371503}}
2022/10/29 01:54:36 lookup(output.txt): FH 0x0100070280e3a60100000000fad64994ded0c36100000000000000000580a6014bdc6b010380a6014adc6b01, attr: {Type:1 FileMode:511 Nlink:1 UID:0 GID:0 Filesize:1916 Used:4096 SpecData:[0 0] FSID:7044703896526771962 Fileid:27688965 Atime:{Seconds:1667022876 Nseconds:651371503} Mtime:{Seconds:1667022876 Nseconds:653371540} Ctime:{Seconds:1667022876 Nseconds:653371540}}
2022/10/29 01:54:36 read(0100070280e3a60100000000fad64994ded0c36100000000000000000580a6014bdc6b010380a6014adc6b01) len=512 offset=0
2022/10/29 01:54:36 read(0100070280e3a60100000000fad64994ded0c36100000000000000000580a6014bdc6b010380a6014adc6b01) len=512 offset=512
2022/10/29 01:54:36 read(0100070280e3a60100000000fad64994ded0c36100000000000000000580a6014bdc6b010380a6014adc6b01) len=256 offset=1024
2022/10/29 01:54:36 read(0100070280e3a60100000000fad64994ded0c36100000000000000000580a6014bdc6b010380a6014adc6b01) len=512 offset=1280
2022/10/29 01:54:36 read(0100070280e3a60100000000fad64994ded0c36100000000000000000580a6014bdc6b010380a6014adc6b01) len=512 offset=1792
2022/10/29 01:54:36 Sums match 3a1b668f5932bc0726c525eeed8b4eb9d5950ca2787652a34ac9dfc5b517c3d1 3a1b668f5932bc0726c525eeed8b4eb9d5950ca2787652a34ac9dfc5b517c3d1
2022/10/29 01:54:36 Completed tests
```

### download

```
PS C:\temp> .\nfsclient.exe 192.168.1.100:/export/share root:0:0 down ./redteam/output.txt
2022/10/29 01:55:18 5
[C:\temp\nfsclient.exe 192.168.1.100:/export/share root:0:0 down ./redteam/output.txt]
2022/10/29 01:55:18 host=192.168.1.100 target=/export/share command=down
2022/10/29 01:55:18 Connecting to 192.168.1.100:892 from unprivileged port
2022/10/29 01:55:18 Connecting to 192.168.1.100:2049 from unprivileged port
2022/10/29 01:55:18 192.168.1.100:/export/share fsinfo=&nfs.FSInfo{Attr:nfs.PostOpAttr{IsSet:false, Attr:nfs.Fattr{Type:0x0, FileMode:0x0, Nlink:0x0, UID:0x0, GID:0x0, Filesize:0x0, Used:0x0, SpecData:[2]uint32{0x0, 0x0}, FSID:0x0, Fileid:0x0, Atime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Mtime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Ctime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}}}, RTMax:0x20000, RTPref:0x20000, RTMult:0x1000, WTMax:0x20000, WTPref:0x20000, WTMult:0x1000, DTPref:0x1000, Size:0xffffffff000, TimeDelta:nfs.NFS3Time{Seconds:0x1, Nseconds:0x0}, Properties:0x1b}
2022/10/29 01:55:18 lookup(redteam): FH 0x0100070180e3a60100000000fad64994ded0c36100000000000000000380a6014adc6b01, attr: {Type:2 FileMode:511 Nlink:2 UID:0 GID:0 Filesize:4096 Used:4096 SpecData:[0 0] FSID:7044703896526771962 Fileid:27688963 Atime:{Seconds:1667022834 Nseconds:221578598} Mtime:{Seconds:1667022876 Nseconds:651371503} Ctime:{Seconds:1667022876 Nseconds:651371503}}
2022/10/29 01:55:18 lookup(output.txt): FH 0x0100070280e3a60100000000fad64994ded0c36100000000000000000580a6014bdc6b010380a6014adc6b01, attr: {Type:1 FileMode:511 Nlink:1 UID:0 GID:0 Filesize:1916 Used:4096 SpecData:[0 0] FSID:7044703896526771962 Fileid:27688965 Atime:{Seconds:1667022876 Nseconds:659371652} Mtime:{Seconds:1667022876 Nseconds:653371540} Ctime:{Seconds:1667022876 Nseconds:653371540}}
2022/10/29 01:55:18 lookup(redteam): FH 0x0100070180e3a60100000000fad64994ded0c36100000000000000000380a6014adc6b01, attr: {Type:2 FileMode:511 Nlink:2 UID:0 GID:0 Filesize:4096 Used:4096 SpecData:[0 0] FSID:7044703896526771962 Fileid:27688963 Atime:{Seconds:1667022834 Nseconds:221578598} Mtime:{Seconds:1667022876 Nseconds:651371503} Ctime:{Seconds:1667022876 Nseconds:651371503}}
2022/10/29 01:55:18 lookup(output.txt): FH 0x0100070280e3a60100000000fad64994ded0c36100000000000000000580a6014bdc6b010380a6014adc6b01, attr: {Type:1 FileMode:511 Nlink:1 UID:0 GID:0 Filesize:1916 Used:4096 SpecData:[0 0] FSID:7044703896526771962 Fileid:27688965 Atime:{Seconds:1667022876 Nseconds:659371652} Mtime:{Seconds:1667022876 Nseconds:653371540} Ctime:{Seconds:1667022876 Nseconds:653371540}}
2022/10/29 01:55:18 FIle size: 1916
2022/10/29 01:55:18 Name: output.txt
2022/10/29 01:55:18 read(0100070280e3a60100000000fad64994ded0c36100000000000000000580a6014bdc6b010380a6014adc6b01) len=1916 offset=0
2022/10/29 01:55:18 Sums match 3a1b668f5932bc0726c525eeed8b4eb9d5950ca2787652a34ac9dfc5b517c3d1 3a1b668f5932bc0726c525eeed8b4eb9d5950ca2787652a34ac9dfc5b517c3d1
2022/10/29 01:55:18 Completed tests
```

### rm

```
PS C:\temp> .\nfsclient.exe 192.168.1.100:/export/share root:0:0 rm ./redteam/output.txt
2022/10/29 01:55:42 5
[C:\temp\nfsclient.exe 192.168.1.100:/export/share root:0:0 rm ./redteam/output.txt]
2022/10/29 01:55:42 host=192.168.1.100 target=/export/share command=rm
2022/10/29 01:55:42 Connecting to 192.168.1.100:892 from unprivileged port
2022/10/29 01:55:42 Connecting to 192.168.1.100:2049 from unprivileged port
2022/10/29 01:55:42 192.168.1.100:/export/share fsinfo=&nfs.FSInfo{Attr:nfs.PostOpAttr{IsSet:false, Attr:nfs.Fattr{Type:0x0, FileMode:0x0, Nlink:0x0, UID:0x0, GID:0x0, Filesize:0x0, Used:0x0, SpecData:[2]uint32{0x0, 0x0}, FSID:0x0, Fileid:0x0, Atime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Mtime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Ctime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}}}, RTMax:0x20000, RTPref:0x20000, RTMult:0x1000, WTMax:0x20000, WTPref:0x20000, WTMult:0x1000, DTPref:0x1000, Size:0xffffffff000, TimeDelta:nfs.NFS3Time{Seconds:0x1, Nseconds:0x0}, Properties:0x1b}
2022/10/29 01:55:42 lookup(redteam): FH 0x0100070180e3a60100000000fad64994ded0c36100000000000000000380a6014adc6b01, attr: {Type:2 FileMode:511 Nlink:2 UID:0 GID:0 Filesize:4096 Used:4096 SpecData:[0 0] FSID:7044703896526771962 Fileid:27688963 Atime:{Seconds:1667022834 Nseconds:221578598} Mtime:{Seconds:1667022876 Nseconds:651371503} Ctime:{Seconds:1667022876 Nseconds:651371503}}
2022/10/29 01:55:42 Completed tests
```

### mkdir

```
PS C:\temp> .\nfsclient.exe 192.168.1.100:/export/share root:0:0 mkdir ./redteam
2022/10/29 01:53:54 5
[C:\temp\nfsclient.exe 192.168.1.100:/export/share root:0:0 mkdir ./redteam]
2022/10/29 01:53:54 host=192.168.1.100 target=/export/share command=mkdir
2022/10/29 01:53:54 Connecting to 192.168.1.100:892 from unprivileged port
2022/10/29 01:53:54 Connecting to 192.168.1.100:2049 from unprivileged port
2022/10/29 01:53:54 192.168.1.100:/export/share fsinfo=&nfs.FSInfo{Attr:nfs.PostOpAttr{IsSet:false, Attr:nfs.Fattr{Type:0x0, FileMode:0x0, Nlink:0x0, UID:0x0, GID:0x0, Filesize:0x0, Used:0x0, SpecData:[2]uint32{0x0, 0x0}, FSID:0x0, Fileid:0x0, Atime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Mtime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Ctime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}}}, RTMax:0x20000, RTPref:0x20000, RTMult:0x1000, WTMax:0x20000, WTPref:0x20000, WTMult:0x1000, DTPref:0x1000, Size:0xffffffff000, TimeDelta:nfs.NFS3Time{Seconds:0x1, Nseconds:0x0}, Properties:0x1b}
2022/10/29 01:53:54 root -> 0x0100070080e3a60100000000fad64994ded0c3610000000000000000
2022/10/29 01:53:54 mkdir(./redteam): created successfully (0x0100070080e3a60100000000fad64994ded0c3610000000000000000)
2022/10/29 01:53:54 Completed tests
```

### rmdir

```
PS C:\temp> .\nfsclient.exe 192.168.1.100:/export/share root:0:0 rmdir ./redteam
2022/10/29 01:56:20 5
[C:\temp\nfsclient.exe 192.168.1.100:/export/share root:0:0 rmdir ./redteam]
2022/10/29 01:56:20 host=192.168.1.100 target=/export/share command=rmdir
2022/10/29 01:56:20 Connecting to 192.168.1.100:892 from unprivileged port
2022/10/29 01:56:20 Connecting to 192.168.1.100:2049 from unprivileged port
2022/10/29 01:56:20 192.168.1.100:/export/share fsinfo=&nfs.FSInfo{Attr:nfs.PostOpAttr{IsSet:false, Attr:nfs.Fattr{Type:0x0, FileMode:0x0, Nlink:0x0, UID:0x0, GID:0x0, Filesize:0x0, Used:0x0, SpecData:[2]uint32{0x0, 0x0}, FSID:0x0, Fileid:0x0, Atime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Mtime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}, Ctime:nfs.NFS3Time{Seconds:0x0, Nseconds:0x0}}}, RTMax:0x20000, RTPref:0x20000, RTMult:0x1000, WTMax:0x20000, WTPref:0x20000, WTMult:0x1000, DTPref:0x1000, Size:0xffffffff000, TimeDelta:nfs.NFS3Time{Seconds:0x1, Nseconds:0x0}, Properties:0x1b}
2022/10/29 01:56:20 root -> 0x0100070080e3a60100000000fad64994ded0c3610000000000000000
2022/10/29 01:56:20 rmdir(redteam): deleted successfully
2022/10/29 01:56:20 Completed tests
```
