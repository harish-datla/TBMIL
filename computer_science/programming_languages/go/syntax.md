
> 💡 Much of the notes here is hand typed / modified based on https://go.dev/tour

### Basics

#### Packages
  - Every go program is made up of packages.
  - programs start running in package `main`.
  - By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package files that being with the statement `package rand`

  - **Example**
    - [View on Github] (https://github.com/harish-datla/TBMIL/blob/main/computer_science/programming_languages/go/code_samples/basics/packages/main.go)
    - To execute on github Codespaces.
      - go to computer_science/programming_languages/go/code_samples/basics/packages directory in terminal and run the run.sh file by typing ./run.sh

#### Exported names
  - In Go, a name can be imported from a package if it begins with a capital letter. For example, Pizza is an exported name, while pizza is not  
  - Similarly, when you are writing packages, any name(variables, structs etc) can only be exported to other packages if its starts with a capital letter.
  - **Example**
    - [View on Github] (https://github.com/harish-datla/TBMIL/blob/main/computer_science/programming_languages/go/code_samples/basics/exported_names/main.go)
    - To execute on github Codespaces.
      - go to computer_science/programming_languages/go/code_samples/basics/exported_names/main.go directory in terminal and run the run.sh file by typing ./run.sh

#### Functions
  - A function can take zero or more arguments.
  - As discussed in exported names section, same rules apply here as well.
  - We will come to this bit later, but if you observe in the example given below, the declaration order is a little bit different than what you are used if you come from c, c++, java background(commonly known as the c family)
  
