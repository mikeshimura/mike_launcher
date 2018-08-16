# mike_launcher
[日本語の説明](https://github.com/mikeshimura/mike_launcher/wiki/%E6%97%A5%E6%9C%AC%E8%AA%9E%E3%81%AE%E8%AA%AC%E6%98%8E)

You pass user only 2 files for install.

For example your application is javadesktop.

$javadesktop.ini and javadesktop.exe for windows or javadesktop for Mac.

Zip and other program data are stored in Amazon S3.

At initial startup, zip file is downloaded and unzipped, and then all files these may change will be downloaded.

After that, godesktop.exe or godesktop execute command which stored in $godesktop.ini.

When program start next time, godesktop.exe or godesktop check S3 repository for update file and
if files are updated, it download them and execute command.

$godesktop.ini content are as follows.

```
[default]
OS = WIN
REGION = us-east-1
AWS_ACCESS_KEY_ID = XXXXXXXXXXXXXXX
AWS_SECRET_ACCESS_KEY = XXXXXXXXXXXXXXXXXXXX
BUCKET =desktoptool
ZIP = $javadesktop.zip
WATCH= $javadesktop-watch.txt
HIS = $his.json
UNZIP = classes.zip
CMD = java -cp lib/*;classes com.mssoftech.javadesktop.Application
```
OS = WIN or MAC

HIS = History file name.

UNZIP = classes.zip   This mean after download classes.zip will be unzipped.

Make sure that AWS_ACCESS_KEY_ID is IAM which has only authorizathion to read S3.

$godesktop-watch.txt content are as follows.

```
assets/tag/index.tag
assets/tag/tagcommon.js
godesktopwin.exe
```

Please down load mike_launcher.exe or mike_launcher from following Google Drive and rename it to your application name.

https://drive.google.com/open?id=1tIh_Ye-6uCAvrXBI7OWG_L7LFq_2ukKs
