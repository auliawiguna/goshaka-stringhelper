Go string helpers inspered by Laravel's :)
# Requirements
- Go version 1.19

# Setup
## Clone the project

- Clone this project by executing `git clone https://github.com/auliawiguna/goshaka-stringhelper`
- Execute `go mod download && go mod verify`

## Usage in your projects
- Note that you need to include the v in the version tag. Execute `go get github.com/auliawiguna/goshaka-stringhelper@v0.1.1`



# Functions

### After

To get string after given mark

```
After(source, mark string) string
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |
| `mark` | `string` | **Required**. Mark text |

Usage:

```
After("This is World of Oz", "is World")
```

Result: ` of Oz`
### AfterLast

To get string after last occurence of given mark

```
  AfterLast(source, mark string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |
| `mark` | `string` | **Required**. Mark text |

Usage:

```
AfterLast("This is World of Oz and World of Colony and the World of Irony", "World of")
```

Result: ` Irony`

### Before

To get string before given mark

```
  Before(source, mark string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |
| `mark` | `string` | **Required**. Mark text |

Usage:

```
Before("This is World of Oz", "is World")
```

Result: `This `

### BeforeLast

To get string before the last occurence of given mark

```
  BeforeLast(source, mark string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |
| `mark` | `string` | **Required**. Mark text |

Usage:

```
BeforeLast("This is World of Oz and World of Colony and the World of Irony", "World of")
```

Result: `This is World of Oz and World of Colony and the  `

### Between

To get string between the first occurence of "start" and the last occurence of "stop"

```
  Between(source, start, stop string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |
| `start` | `string` | **Required**. Mark text |
| `stop` | `string` | **Required**. Mark text |

Usage:

```
Between("This is World of Oz", "This", "of")
```

Result: ` is World `

### BetweenFirst

To get the smallest possible portion of a string between "start" and "stop"

```
  BetweenFirst(source, start, stop string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |
| `start` | `string` | **Required**. Mark text |
| `stop` | `string` | **Required**. Mark text |

Usage:

```
BetweenFirst("[a] [b] XX [c]", "[", "]")
```

Result: `a`

### Camel

To convert str into camel case (camelCase)

```
  Camel(source string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |

Usage:

```
Camel("i feel good")
```

Result: `iFeelGood`

### Pascal

To convert str into pascal case (PascalCase)

```
  Pascal(source string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |

Usage:

```
Pascal("i feel good")
```

Result: `IFeelGood`

### Snake

To convert str into snake case (snake_case)

```
  Snake(source string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |

Usage:

```
Snake("i feel good")
```

Result: `i_feel_good`

### Kebab

To convert str into kebab case (kebab-case)

```
  Kebab(source string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |

Usage:

```
Kebab("i feel good")
```

Result: `i-feel-good`

### Headline

To convert str into headline case (Headline Case)

```
  Headline(source string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `source` | `string` | **Required**. Source text |

Usage:

```
Headline("i feel good")
```

Result: `I Feel Good`

### Contains

To determine if target is exists in str

```
  Contains(str, target string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `str` | `string` | **Required**. Source text |
| `target` | `string` | **Required**. Target text |

Usage:

```
Contains("i feel good", "feel")
```

### ContainsAll

To determine if array of target are/is exists in str

```
  ContainsAll(str string, target []string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `str` | `string` | **Required**. Source text |
| `target` | `[]string` | **Required**. Array of the target text |

Usage:

```
ContainsAll("i feel good", []string{"i", "feel", "good"})
```

Result: `true`

### EndsWith

To determine if str is ends with target

```
  EndsWith(str string, target {}interface)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `str` | `string` | **Required**. Source text |
| `target` | `array|string` | **Required**. Target text |

Usage:

```
EndsWith("This Is Title", "Not")
```

Result: `false`

### Finish

To adds a "add" to "str" if it does not already end with "add"

```
  Finish(str, add string)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `str` | `string` | **Required**. Source text |
| `add` | `string` | **Required**. Target text |

Usage:

```
Finish("This Is Title", "Title")
```

Result: `This Is Title`

# Running Tests

To run tests, run the following command

```bash
  go test
```