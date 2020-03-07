# backup-settings

This program backups chosen directories and files as configured in a `backup-config.json` JSON file put in any folder within the executing folder

## Configuration files

```json
{
    "user_path": "c:\\path\\of\\user",
    "folders_to_save": [
        ".ssh",
        ".android",
        ".blabla"
    ],
    "folders_to_ignore": [
        ".ssh/archive",
        ".android/avd"
    ]
}
```

## Todos

- [x] works with `/` or `\` in path
- [ ] improve backup (e.g. backup elsewhere than in this repo)
- [ ] allow reverse copy of backup to system (e.g. for a quick reinstall of the system)
