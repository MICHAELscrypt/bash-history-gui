## Tables
```sql
CREATE TABLE "applications" (
	"application"	TEXT NOT NULL UNIQUE,
	"occurrences"	INTEGER NOT NULL,
	"display"	TEXT DEFAULT 'True',
	"last_used"	INTEGER,
	"favorite"	TEXT DEFAULT 'False'
);
```

```sql
CREATE TABLE "commands" (
	"application"	TEXT NOT NULL,
	"full_command"	TEXT NOT NULL UNIQUE,
	"deleted"	TEXT DEFAULT 'False',
	"last_used"	INTEGER,
	"favorite"	TEXT DEFAULT 'False'
);
```

## Triggers

```sql
CREATE TRIGGER update_occurrences_after_insert
AFTER INSERT ON commands
BEGIN
    UPDATE applications
    SET occurrences = occurrences + 1
    WHERE application = NEW.application;
END;
```

```sql
CREATE TRIGGER update_occurrences_after_delete
AFTER DELETE ON commands
BEGIN
    UPDATE applications
    SET occurrences = occurrences - 1
    WHERE application = OLD.application;
END;
```
