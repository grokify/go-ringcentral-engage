package main


var map[string]string{
	"getProductActiveCallsListRequest": ```    @RequestMapping(path = "activeCalls/list", method = RequestMethod.GET, produces = JSON_MEDIA_TYPE)
    @ApiOperation(value = "Returns a listing of current active calls for a given product on an account", notes = "Permissions: READ on Account")
    @PreAuthorize("hasPermission('ACCT', 'ACCT')")
    public List<ActiveCallListResponse> getProductActiveCallsListRequest(
            @PathVariable("accountId") final String accountId,
            @RequestParam(required = true) final Product product,
            @RequestParam(required = true) final Integer productId,
            @RequestParam(name = "page", required = false) final Integer page,
            @RequestParam(name = "maxRows", required = false) final Integer maxRows) {```}

/*
			Reddit Source
			HubSpot submited.*/

			func main () {
				fmt.Println("DONE")
			}