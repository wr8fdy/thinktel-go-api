generate:
    rm -rf client models
    wget -O swagger.json https://api.thinktel.ca/swagger/ui/swagger.json
    # workaround to avoid di_ds_sip_trunk names
    sed -i 's/DIDs (SIP Trunks)/dids/g' swagger.json
    /usr/bin/swagger generate client --additional-initialism=DID --additional-initialism=SIP --name=thinktel-ucontrol
