package example

func hoge() {
	json := ""

	// 正しくフォーマットされること
	json = `
{
  "key1": "value1",
  "key2": {
    "key3": "value3"
  }
}`

	// インデントがおかしいJSONも正しくフォーマットされること
	json = `
{
  "key1": "value1",
  "key2": {
    "key3": "value3"
  }
}`

	// フィールドの順番が維持されること
	json = `
{
  "b": "value1",
  "a": "value2"
}`

	// 一部の記号がUnicodeエスケープされないこと
	json = `
{
  "key": "< > &",
  "dummy": "dummy"
}`

	// Unicodeエスケープはそのまま維持されること
	json = `
{
  "key": "\u0022 \u005c \u002f \u0009"
}`

	// 20文字以下のJSONはフォーマットされないこと
	json = `{"key1":"value1"}`

	// ダブルクォーテーションで囲まれたJSONはフォーマットされないこと
	json = "{\"key1\":\"value1\",\"key2\":\"value2\"}"

	// 不正なJSONの場合はフォーマットされないこと
	json = `{"key1"}`
	json = `{""}`
	json = `{}`

	// SQLは無視されること
	sql := `SELECT * FROM table WHERE id = 1`
}
