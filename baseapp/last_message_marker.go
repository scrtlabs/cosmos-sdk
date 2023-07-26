package baseapp

type LastMsgMarkerContainer struct {
	marker bool
}

func (mng *LastMsgMarkerContainer) SetMarker(value bool) {
	mng.marker = value
}

func (mng *LastMsgMarkerContainer) GetMarker() bool {
	return mng.marker
}
