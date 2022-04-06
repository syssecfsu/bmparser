## Fuzzing C Implementation

This file contains basic instructions to run AFL++ on the test projects.

Because fuzz testing is resource intensive, it is strongly recommended not to run
AFL++ in a VM unless you have a fast computer. On Windows, you can run AFL++ in
**WSL2**. Follow the instructions from Microsoft to install WSL2. If you are running Linux, make sure it is a recent version of Linux (e.g., Ubuntu 21.10 or Fedora 35). We use WSL2 as an example here.

1. Follow the instructions from Microsoft to install WSL2.

2. Install Ubuntu on Microsoft Store, and follw the instructions [Here](https://askubuntu.com/questions/1369637/how-do-i-install-ubuntu-21-10-in-wsl-on-windows-11) to upgrade Ubuntu from 20.04 to 21.10.

3. Start Ubuntu and run `sudo apt update` and `sudo apt upgrade` to update your Ubuntu system.

4. Run the following command to install basic build utilities.

    `sudo apt-get install -y build-essential python3-dev automake cmake git flex bison libglib2.0-dev libpixman-1-dev python3-setuptools`

5. Install llvm and clang.

    `sudo apt-get install -y lld llvm llvm-dev clang`

6. Install gcc and its dev-lib

    `sudo apt-get install -y gcc-$(gcc --version|head -n1|sed 's/\..*//'|sed 's/.* //')-plugin-dev libstdc++-$(gcc --version|head -n1|sed 's/\..*//'|sed 's/.* //')-dev`

7. Install tools for QEMU build.

    `sudo apt-get install -y ninja-build # for QEMU mode`

8. Clone the AFL repo:

    `git clone https://github.com/AFLplusplus/AFLplusplus`

9. Build and install AFL++

    ```sh
    cd AFLplusplus
    make distrib
    sudo make install
    ```

10. Check `/usr/local/bin`, it should show something similar to the following:

    ```console
    /usr/local/bin$ ls -lh
    total 7.5M
    -rwxr-xr-x 1 root root 284K Apr  5 18:19 afl-analyze
    lrwxrwxrwx 1 root root    6 Apr  5 18:19 afl-c++ -> afl-cc
    -rwxr-xr-x 1 root root 141K Apr  5 18:19 afl-cc
    lrwxrwxrwx 1 root root    6 Apr  5 18:19 afl-clang -> afl-cc
    lrwxrwxrwx 1 root root    6 Apr  5 18:19 afl-clang++ -> afl-cc
    lrwxrwxrwx 1 root root    6 Apr  5 18:19 afl-clang-fast -> afl-cc
    lrwxrwxrwx 1 root root    9 Apr  5 18:19 afl-clang-fast++ -> ./afl-c++
    lrwxrwxrwx 1 root root    6 Apr  5 18:19 afl-clang-lto -> afl-cc
    lrwxrwxrwx 1 root root    9 Apr  5 18:19 afl-clang-lto++ -> ./afl-c++
    -rwxr-xr-x 1 root root  17K Apr  5 18:19 afl-cmin
    -rwxr-xr-x 1 root root  13K Apr  5 18:19 afl-cmin.bash
    -rwxr-xr-x 1 root root 1.5M Apr  5 18:19 afl-fuzz
    lrwxrwxrwx 1 root root    6 Apr  5 18:19 afl-g++ -> afl-cc
    lrwxrwxrwx 1 root root    7 Apr  5 18:19 afl-g++-fast -> afl-c++
    lrwxrwxrwx 1 root root    6 Apr  5 18:19 afl-gcc -> afl-cc
    lrwxrwxrwx 1 root root    6 Apr  5 18:19 afl-gcc-fast -> afl-cc
    -rwxr-xr-x 1 root root  43K Apr  5 18:19 afl-gotcpu
    -rwxr-xr-x 1 root root  31K Apr  5 18:19 afl-ld-lto
    lrwxrwxrwx 1 root root    6 Apr  5 18:19 afl-lto -> afl-cc
    lrwxrwxrwx 1 root root    6 Apr  5 18:19 afl-lto++ -> afl-cc
    -rwxr-xr-x 1 root root  17K Apr  5 18:19 afl-network-client
    -rwxr-xr-x 1 root root 110K Apr  5 18:19 afl-network-server
    -rwxr-xr-x 1 root root 4.3K Apr  5 18:19 afl-persistent-config
    -rwxr-xr-x 1 root root 7.7K Apr  5 18:19 afl-plot
    -rwxr-xr-x 1 root root 4.8M Apr  5 18:19 afl-qemu-trace
    -rwxr-xr-x 1 root root 320K Apr  5 18:19 afl-showmap
    -rwxr-xr-x 1 root root 5.7K Apr  5 18:19 afl-system-config
    -rwxr-xr-x 1 root root 322K Apr  5 18:19 afl-tmin
    -rwxr-xr-x 1 root root 7.5K Apr  5 18:19 afl-whatsup
    ```

    Particularly, you should see `afl-clang-fast` and `afl-clang-lto`.

11. Get Project4_tests.zip file from your canvas email and put it in your Windows file system and unzip it. Add the test files found in this repo (under the `afl++/bmp` directory) to the uncompressed directory.

12. Go to Ubuntu system, copy the uncompressed directory to your home directory. You can access the Windows file system under `/mnt/c` or `/mnt/d` etc.

13. Update the Makefile to the following:

    ```Makefile
    CC=afl-clang-lto
    LD=afl-clang-fast
    CFLAGS=-g -std=c99

    bmp_parser: bmp_parser.o
        $(CC) $(CFLAGS) -o bmp_parser bmp_parser.o -lm

    bmp_parser.o: bmp_parser.c
        $(CC) -c $(CFLAGS) bmp_parser.c

    clean:
        rm -f *.x *.o bmp_parser
    ```

    Note that Makefile use tab for indention, make sure you copied Makefile works.

14. Build the bmp_parser with the `make` command. Enable `ASAN` and `UBSAN` as shown in the AFL++ tutorial. Create a directory to run AFL++ on bmp_parser. Copy bmp_parser to this directory.

    If you enable `ASAN`, you will immediately find an off-by-one bug on line 67. You need to fix it first.

    The clang build will also complain a bug in line 125. You can leave it as is or fix it. Just record in your report.

15. Use the following command to run AFL++ on bmp_parser. If everything is correct, you should get some crashes in a couple of minutes.

    `afl-fuzz -i ../bmp -o out -e bmp -D -- ./bmp_parser @@`

## Fuzzing Python Implementation

To add later.
