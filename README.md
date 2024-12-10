# Avatar Generator

Generates random pixel avatars of any size and any color.

## Usage 

### Print to console

```go
package main

import avatargen "github.com/1N1Group/avatar_gen"

func main() {
	avatar := avatargen.New()
	avatar.Print()
}
```

#### Result
```shell
                        
    ██  ██    ██  ██    
    ██            ██    
  ██    ████████    ██  
          ████          
      ████    ████      
  ██  ██  ████  ██  ██  
  ██  ████████████  ██  
    ██████    ██████    
      ██        ██      
  ████  ████████  ████  
                        
```


### Write to file

```go
package main

import (
	"os"

	avatargen "github.com/1N1Group/avatar_gen"
)

func main() {
	avatar := avatargen.New()
	buff := avatar.ToBuffer(1111)

	file, err := os.Create("img.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buff.WriteTo(file)
}

```

#### Result

![image](img.jpeg)
