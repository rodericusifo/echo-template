# ECHO TEMPLATE'S NOTE

Note for Echo Template

## DECLARING GORM MODELS

- If you want to use field alias (with operator `AS`)

  - Step 1: Declare in Model Struct's Field

    ```properties
    type Something struct {
      // Migrated Fields
      ID        uint   `gorm:"primaryKey"`
      XID       string `gorm:"column:xid"`
      ...
      CreatedAt time.Time
      UpdatedAt time.Time
      DeletedAt gorm.DeletedAt `gorm:"index"`

      // Relations
      ...

      // Not Migrated Fields
      SomethingXID string `gorm:"column:something_xid;<-:false;-:migration"`
      ...
    }
    ```

  - Step 2: Type Select Statement Like This

    ```properties
    ...
    &types.Query{
      Selects: []types.SelectOperation{
        {Field: "id"},
        {Field: "xid", Alias: "something_xid"},
        ...
        {Field: "created_at"},
        {Field: "updated_at"},
      },
    }
    ...
    ```
