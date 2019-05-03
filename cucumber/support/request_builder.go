package support

// ConversionRequest is a representation of data fields in a conversion request body
type ConversionRequest struct {
	ImageLocation string `json:"image_location"`
	DesiredFormat string `json:"desired_format"`
}

// ConversionResponse is a representation of data fields in a conversion response body
type ConversionResponse struct {
	OriginalImage string `json:"original_image"`
	DesiredFormat string `json:"desired_format"`
	NewImage      string `json:"new_image"`
}

// ConversionRequestBody returns a conversion request body based on the information provided
func ConversionRequestBody(imageLocation, desiredFormat string) ConversionRequest {
	return ConversionRequest{
		ImageLocation: imageLocation,
		DesiredFormat: desiredFormat,
	}
}
