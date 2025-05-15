
> 💡 Much of the notes here is hand typed / modified based on https://go.dev/tour

### Basics

#### Packages
  - Every go program is made up of packages.
  - programs start running in package `main`.
  - By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package files that being with the statement `package rand`

  - **Example**
    [View on Github] (https://github.com/harish-datla/TBMIL/blob/main/computer_science/programming_languages/go/code_samples/basics/packages/main.go)
    To execute on github Codespaces.
    Cd to computer_science/programming_languages/go/code_samples/basics/packages and run the run.sh file by typing ./run.sh

#### Exported names
  - In Go, a name can be imported from a package if it begins with a capital letter. For example, Pizza isan exported name  
  - Similarly, when you are writing packages, any name(variables, structs etc) can only be exported to other packages if its starts with a capital letter.
    
