# search-api
For development
===============
1. Clone the repo and open the folder
`cd search-api`
2. Run setup.sh to set up certificates
`source setup.sh`
3. Run the file server.go
`go run server.go`
4. Open the graphql link in browser.
5. Issue a search query in the input.
Example:
```
query{
    search(input: [
        {
            filters: [
                {
                    property: "kind",
                    values:[ "pod" ]
                }
            ],
        }
    ]) {
    		related{
          kind
          count
        }
        count 
    }
}
```
And the output appears on the right hand side.