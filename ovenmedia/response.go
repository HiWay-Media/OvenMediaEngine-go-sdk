package ovenmedia

import "time"

type BaseResponseOK struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

// Response<VirtualHost>
type ResponseVirtualHost struct {
	VRHosts []VirtualHost
}

type VirtualHost struct {
	BaseResponseOK
	Response VRHostResponse `json:"response"`
}

type VRHostResponse struct {
	Name string `json:"name"`
}

type ResponseVirtualList struct {
	BaseResponseOK
	Response []string `json:"response"`
}

type ResponseStreamInfo struct {
	BaseResponseOK
	Response struct {
		CreatedTime time.Time `json:"createdTime"`
		SourceType  string    `json:"sourceType"`
		Tracks      []struct {
			Id    int    `json:"id"`
			Type  string `json:"type"`
			Video struct {
				Bitrate          string  `json:"bitrate"`
				Bypass           bool    `json:"bypass"`
				Codec            string  `json:"codec"`
				Framerate        float64 `json:"framerate"`
				HasBframes       bool    `json:"hasBframes"`
				KeyFrameInterval int     `json:"keyFrameInterval"`
				Height           int     `json:"height"`
				Width            int     `json:"width"`
			} `json:"video,omitempty"`
			Audio struct {
				Bitrate    string `json:"bitrate"`
				Bypass     bool   `json:"bypass"`
				Channel    int    `json:"channel"`
				Codec      string `json:"codec"`
				Samplerate int    `json:"samplerate"`
			} `json:"audio,omitempty"`
		} `json:"tracks"`
	} `json:"input"`
	Name    string `json:"name"`
	Outputs []struct {
		Name   string `json:"name"`
		Tracks []struct {
			Id    int    `json:"id"`
			Type  string `json:"type"`
			Video struct {
				Bypass bool `json:"bypass"`
			} `json:"video,omitempty"`
			Audio struct {
				Bypass     bool   `json:"bypass"`
				Bitrate    string `json:"bitrate,omitempty"`
				Channel    int    `json:"channel,omitempty"`
				Codec      string `json:"codec,omitempty"`
				Samplerate int    `json:"samplerate,omitempty"`
			} `json:"audio,omitempty"`
		} `json:"tracks"`
	} `json:"outputs"`
}

// PUSH Stuff
type ResponseStartPush struct {
	BaseResponseOK
	Response ResponsePush `json:"response"`
}

type ResponsePush struct {
	App         string `json:"app"`
	CreatedTime string `json:"createdTime"`
	ID          string `json:"id"`
	Protocol    string `json:"protocol"`
	SentBytes   int    `json:"sentBytes"`
	SentTime    int    `json:"sentTime"`
	Sequence    int    `json:"sequence"`
	StartTime   string `json:"startTime"`
	State       string `json:"state"`
	Stream      struct {
		Name   string `json:"name"`
		Tracks []int  `json:"tracks"`
	} `json:"stream"`
	StreamKey      string `json:"streamKey"`
	TotalSentBytes int    `json:"totalsentBytes"`
	TotalSentTime  int    `json:"totalsentTime"`
	URL            string `json:"url"`
	Vhost          string `json:"vhost"`
}

type ResponsePushes struct {
	BaseResponseOK
	Response []ResponsePush `json:"response"`
}

/*
	{
		"createdTime": "2021-01-11T02:52:22.013+09:00",
		"lastRecvTime": "2021-01-11T04:11:41.734+09:00",
		"lastSentTime": "2021-01-11T02:52:22.013+09:00",
		"lastUpdatedTime": "2021-01-11T04:11:41.734+09:00",
		"maxTotalConnectionTime": "2021-01-11T02:52:22.013+09:00",
		"maxTotalConnections": 0,
		"totalBytesIn": 494713880,
		"totalBytesOut": 0,
		"totalConnections": 0
	}
*/
type ResponseStats struct {
	CreatedTime            string `json:"createdTime"`
	LastRecvTime           string `json:"lastRecvTime"`
	LastSentTime           string `json:"lastSentTime"`
	LastUpdatedTime        string `json:"lastUpdatedTime"`
	MaxTotalConnectionTime string `json:"maxTotalConnectionTime"`
	MaxTotalConnections    int    `json:"maxTotalConnections"`
	TotalBytesIn           int    `json:"totalBytesIn"`
	TotalBytesOut          int    `json:"totalBytesOut"`
	TotalConnections       int    `json:"totalConnections"`
}

// Recording Stuff

/*
	{
	    "message": "OK",
	    "response": [
	        {
	            "state": "ready",
	            "id": "stream_o",
	            "vhost": "default",
	            "app": "app",
	            "stream": {
	                "name": "stream_o",
	                "tracks": []
	            },
	            "filePath": "/path/to/save/recorded/file_${Sequence}.ts",
	            "infoPath": "/path/to/save/information/file.xml",
	            "interval": 60000,
	            "schedule": "0 0 *1",
	            "segmentationRule": "continuity",
	            "createdTime": "2021-08-31T23:44:44.789+0900"
	        }
	    ],
	    "statusCode": 200
	}
*/
type ResponseRecordingStart struct {
	BaseResponseOK
	Response []ResponseRecording `json:"response"`
}

type ResponseRecording struct {
	State  string `json:"state"`
	ID     string `json:"id"`
	Vhost  string `json:"vhost"`
	App    string `json:"app"`
	Stream struct {
		Name   string `json:"name"`
		Tracks []int  `json:"tracks"`
	} `json:"stream"`
	FilePath         string `json:"filePath"`
	InfoPath         string `json:"infoPath"`
	Interval         int    `json:"interval"`
	Schedule         string `json:"schedule"`
	SegmentationRule string `json:"segmentationRule"`
	CreatedTime      string `json:"createdTime"`
}

//
