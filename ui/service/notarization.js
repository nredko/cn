const NotarizationService = {
  fetch: function (query, callback) {
    const xmlHttpRequest = new XMLHttpRequest();
    xmlHttpRequest.onreadystatechange = function () {
      if (this.readyState === 4 && this.status === 200) {
        callback(JSON.parse(this.responseText));
      }
    };
    xmlHttpRequest.open("GET", "/containers?query=" + query, true);
    xmlHttpRequest.send();
  },
  history: function (hash, callback) {
    const xmlHttpRequest = new XMLHttpRequest();
    xmlHttpRequest.onreadystatechange = function () {
      if (this.readyState === 4 && this.status === 200) {
        callback(JSON.parse(this.responseText));
      }
    };
    xmlHttpRequest.open("GET", "/history?hash=" + hash, true);
    xmlHttpRequest.send();
  },
  notarize: function (hash, status, callback) {
    const xmlHttpRequest = new XMLHttpRequest();
    xmlHttpRequest.onreadystatechange = callback;
    xmlHttpRequest.open("POST", "/notarize?hash=" + hash + '&status=' + status, true);
    xmlHttpRequest.send();
  },
  bulkNotarize: function (query, callback) {
    const xmlHttpRequest = new XMLHttpRequest();
    xmlHttpRequest.onreadystatechange = callback;
    xmlHttpRequest.open("POST", "/bulk-notarize?status=Notarized&query=" + query, true);
    xmlHttpRequest.send();
  },
  bulkUnTrust: function (query, callback) {
    const xmlHttpRequest = new XMLHttpRequest();
    xmlHttpRequest.onreadystatechange = callback;
    xmlHttpRequest.open("POST", "/bulk-notarize?status=Untrusted&query=" + query, true);
    xmlHttpRequest.send();
  }
};

export {NotarizationService};
