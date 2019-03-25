const en_US = {
    'app.tasks': 'TASKS',
    'tasklist.title': 'TASKS',
    'tasklist.new': 'new +',
    'tasklist.table.protocol': 'Protocol',
    'tasklist.table.Id': 'ID',
    'tasklist.table.host': 'Host',
    'tasklist.table.createdTime': 'CreatedTime',
    'tasklist.table.updatedTime': 'UpdatedTime',
    'tasklist.table.operation': 'Operation',
    'tasklist.table.start': 'start',
    'tasklist.table.delete': 'delete',
    'tasklist.table.stop': 'stop',
    'tasklist.table.status': 'Status',
    'tasklist.table.result': 'Result',
    'tasklist.table.resume': 'resume',
    'tasklist.table.running': 'running',
    'tasklist.table.fail': 'fail',
    'tasklist.table.success': 'success',
    'tasklist.table.unfinished': 'unfinished',


    'taskmodal.title': 'New Task',

    'taskform.nickname.message':'please inpu your nickname',
    'taskform.protocol.label': 'protocol',
    'taskform.protocol.placeholder': 'please select protocol',
    'taskform.URL.label':'URL',
    'taskform.URL.placeholder':'please input URL',
    'taskform.Requests.label':'Requests',
    'taskform.Requests.placeholder':'please input Requests count',
    'taskform.Method.label':'Method',
    'taskform.Method.placeholder':'please select Method',
    'taskform.Concurrency.label':'Concurrency',
    'taskform.Concurrency.placeholder':'please input Concurrency count',
    'taskform.Timeout.label':'Timeout',
    'taskform.Timeout.placeholder':'please input timeout(seconds)',
    'taskform.Duration.label':'Duration',
    'taskform.Duration.placeholder':'please input Duration(seconds)',
    'taskform.ContentType.label':'ContentType',
    'taskform.ContentType.placeholder':'please select ContentType',
    'taskform.DisableKeepAlive.label':'DisableKeepAlive',
    
    //websocket options
    'taskform.Scheme.placeholder':'please input Scheme',
    'taskform.Scheme.label':'Scheme',
    'taskform.Host.placeholder':'please input Host',
    'taskform.Host.label':'Host',
    'taskform.Path.placeholder':'please input Path',
    'taskform.Path.label':'Path',
    'taskform.Body.placeholder':'please input Body',
    'taskform.Body.label':'Body',
    
}
export default en_US;

// "URL":              "http://127.0.0.1:8080",
//         "Requests":         0,
//         "Method":           "GET",
//         "Concurrency":      5,
//         "Timeout":          30,
//         "Duration":         10,
//         "ContentType":      "text/plain",
//         "MaxResults":       1000000,
//         "DisableKeepAlive": true

// Scheme      string      `json:"scheme,omitempty"`
// 	Host        string      `json:"host,omitempty"`
// 	Path        string      `json:"path,omitempty"`
// 	Requests    int         `json:"requests,omitempty"`
// 	Concurrency int         `json:"concurrency,omitempty"`
// 	Timeout     int         `json:"timeout,omitempty"`
// 	Duration    int         `json:"duration,omitempty"`
// 	Body        string      `json:"body,omitempty"`
// 	MaxResults  int         `json:"maxResults,omitempty"`