# backup-settings

This program backups chosen settings as configured in a `backup-config.json` file put in a folder within the executing folder

## Configuration files

```json
{
    "user_path": "c:\\path\\of\\user",
    "folders_to_save": [
        ".ssh",
        ".android",
        ".blabla"
    ]
}
```

## Todos

- [ ] improve backup (e.g. backup elsewhere than in this repo)
- [ ] allow reverse copy of backup to system (e.g. for a quick reinstall of the system)
