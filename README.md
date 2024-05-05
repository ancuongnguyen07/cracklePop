# Cracklepop
The CLI tool which prints out the numbers from i to j (inclusive). If the number is divisible by 3,
print Crackle instead of the number. If it's divisible by 5, print Pop. If it's divisible
by both 3 and 5, print CracklePop

## Build
```sh
cd cracklePop
make
```

## Run
```sh
cd build
./cracklepop print --start 1 --end 100
```

Or run the following to show helping messages:
```sh
./cracklepop
```