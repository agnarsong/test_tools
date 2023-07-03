#/bin/bash

if [ -z "$1" ]; then
    echo "yarn add..."
    cd packages/contracts
    yarn
    yarn hardhat clean
    yarn hardhat compile
    echo "hardhat compiled\n\n"

    echo "deploy contracts..."
    yarn hardhat run scripts/deploy.js --network l2
    echo "deploy contracts compiled"
    cd -
fi

cd packages/contracts
rm -rf build

for file in `ls artifacts/contracts`
    do 
        path="artifacts/contracts/$file"
        if [[ -d $path ]]; then
            subFile="${file/sol/json}"
            bytecode=`cat "$path/$subFile" |jq -r ".bytecode"` 
        fi
        
        if [ $bytecode != "0x" ]; then
            package="${file/.sol}"
            solFile="contracts/$file"
            abiFile="contracts_${package}_sol_${package}.abi"

            yarn solc $solFile --abi -o build --base-path './' --include-path 'node_modules'
            # solc --bin $solFile -o build --base-path './' --include-path 'node_modules'

            if [[ ! -d "../../contracts/$package" ]]; then
                mkdir "../../contracts/$package"
            fi

            # abigen --abi=build/$package.abi --pkg=$package --out=../../contracts/$package/$package.go --bin=build/$package.bin
            abigen --abi=build/$abiFile --pkg=$package --out=../../contracts/$package/$package.go
        fi

    done

cd -