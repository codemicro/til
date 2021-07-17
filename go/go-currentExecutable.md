# Get the path and name of the currently running executable

There are two ways to accomplish this.

 * Use command line arguments to get the name used to invoke the program
   ```go
   os.Args[0] // -> "/tmp/___go_build_github_com_codemicro_cligen_testdata_package"
   ```
 * Use `os.Executable`
   ```go
   exec, err := os.Executable()
   if err != nil {
       // do a thing
   }
   exec // -> "/tmpfs/play"
   ```
