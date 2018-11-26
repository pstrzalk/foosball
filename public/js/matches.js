function viewHandler() {
  var pullParams = {
    FunctionName: 'GetMatches',
    InvocationType: 'RequestResponse',
    LogType: 'None',
    Payload: JSON.stringify({ per_page: 50 })
  };

  var rowTemplate = function(data) {
    var html = '';
    var score_1 = data.scores[0];
    var score_2 = data.scores[1];

    var t = `<tr>
      <td>${score_1.players[0].name} & ${score_1.players[1].name}</td>
      <td>${score_1.score}</td>
      <td>${score_2.players[0].name} & ${score_2.players[1].name}</td>
      <td>${score_2.score}</td>
    </tr>`;

    return t;
  };

  var tableTemplate = function(results) {
    return `<table>${results.map(data => rowTemplate(data)).join('')}</table>`;
  };

  var lambdaHandler = function(err, data) {
    if (err) {
      console.warn(err);
    } else {
      pullResults = JSON.parse(JSON.parse(data.Payload));
      document.getElementById('matches').innerHTML = tableTemplate(pullResults);
    }
  }

  lambda.invoke(pullParams, lambdaHandler);
}

window.addEventListener('load', viewHandler);
