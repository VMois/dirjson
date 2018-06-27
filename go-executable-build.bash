# Origin: https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
package=$1
custom_output_name=$2
if [ -z "$package" ] || [ -z "$custom_output_name" ]; then
  echo "usage: $0 <package-name> <custom_output_name>"
  exit 1
fi

package_split=(${package//\// })
package_name=${package_split[-1]}

platforms=("windows/amd64" "linux/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$custom_output_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi  

    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o $output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done