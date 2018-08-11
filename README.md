# mike_launcher
[日本語の説明](https://github.com/mikeshimura/mike_launcher/wiki/%E6%97%A5%E6%9C%AC%E8%AA%9E%E3%81%AE%E8%AA%AC%E6%98%8E)

You pass user only 2 files for install.

For example your application is godesktop.

$godesktop.ini and godesktop.exe for windows or godesktop for Mac.

Zip and other program data are stored in Amazon S3.

At initial startup zip file is downloaded and unzipped, and then all files these may change will be downloaded.

After that, godesktop.exe or godesktop execute command which stored in $godesktop.ini.

When program start next time, godesktop.exe or godesktop check S3 repository for update file and
if files are updated, it download them and execute command.

$godesktop.ini content are as follows.

```
[default]
OS = WIN
AWS_ACCESS_KEY_ID = AKIAJ5DFKKOTLS2J4IFA
AWS_SECRET_ACCESS_KEY = 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
BUCKET =desktoptool
ZIP = $godesktop.zip
WATCH= $godesktop-watch.txt
HIS = $his.json
CMD = godesktopwin.exe
```
