# 大学の課題を通知するdiscord BOT

## Require
### Env
You can run in test mode when you add "-t" option.
  
In test mode, must locate `.env` in the same dir as app.
```
DISCORD_TOKEN=example
CHANNEL_ID=example
STUDENT_ID=example
PASSWORD=example
```

## Settings
### Exclude specify courses
You can select courses which this app don't need to notify.
  
You use `--exclude {filepath}` option to exclude.
```
your_course_name_1
your_course_name_2
your_course_name_3
```
