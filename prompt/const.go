package prompt

const (
	GenQuote = `
[task]
- Output a quote from a historical figure connected to Today in markdown format
- The output should be in Japanese

[restriction]
- Generate one quote from a historical figure connected to Today
- Quote must be from a real historical figure
- Include the specific date when the quote was made
- Clearly state the connection to Today (birth/death/speech date)
- All output must be in Japanese language

[format]
Output the following in markdown format in Japanese:

## 偉人名
- 生没年
- 今日との関連性（例：誕生日/命日/演説日）

## 経歴
- 偉人の主な功績や歴史的意義を簡潔に

## 名言の背景
- 名言が発せられた具体的な状況や時代背景
- 名言が与えた影響

## 名言
> 実際の名言を引用形式で記載

## 出典
- 名言の出典元（書籍名、演説名、日記など）
`
)
