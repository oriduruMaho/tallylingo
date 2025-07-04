# show help message
@default: help

App := 'TallyLingo'
Version := `grep '^const VERSION = ' cmd/main/version.go | sed "s/^VERSION = \"\(.*\)\"/\1/g"`

# show help message
@help:
    echo "Build tool for {{ App }} {{ Version }} with Just"
    echo "Usage: just <recipe>"
    echo ""
    just --list

# build the application
build: test
    go build -o tallylingo cmd/main

# run tests and generate coverage report
test:
    go test -covermode=count -coverprofile=coverage.out ./...

# clean up build artifacts
clean:
    go clean
    rm -f tallylingo coverage.out build

# update version if the new version is provided
update_version new_version = "":
    if [ "{{ new_version }}" != "" ]; then \
        sed 's/$VERSION/{{ new_version }}/g' .template/README.md > README.md; \
        sed 's/$VERSION/{{ new_version }}/g' .template/version.go > cmd/main/version.go; \
    fi

# build tallylingo for all platforms
make_distribution_files:
    for os in "linux" "windows" "darwin"; do \
        for arch in "amd64" "arm64"; do \
            mkdir -p dist/{{ App }}-$os-$arch; \
            env GOOS=$os GOARCH=$arch go build -o dist/{{ App }}-$os-$arch/{{ App }} cmd/main/{{ App }}.go; \
            cp README.md LICENSE dist/{{ App }}-$os-$arch; \
            tar cvfz dist/t{{ App }}-$os-$arch.tar.gz -C dist {{ App }}-$os-$arch; \
        done; \
    done

upload_assets tag:
    gh release upload --repo oriduruMaho/{{ App }} {{ tag }} dist/*.tar.gz
