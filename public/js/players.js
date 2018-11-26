function viewHandler() {
  var indexPullParams = {
    FunctionName : 'GetPlayers',
    InvocationType : 'RequestResponse',
    LogType : 'None',
  };

  var createPullParams = {
    FunctionName : 'CreatePlayer',
    InvocationType : 'RequestResponse',
    LogType : 'None',
  };

  var rowTemplate = function(player) {
    return `<tr>
      <td>${player.id}</td>
      <td>${player.name}</td>
    </tr>`;
  };

  var tableTemplate = function(results) {
    return `<table>${results.map(data => rowTemplate(data)).join('')}</table>`;
  };

  var lambdaIndexHandler = function(err, data) {
    if (err) {
      console.warn(err);
    } else {
      pullResults = JSON.parse(JSON.parse(data.Payload));
      document.getElementById('players').innerHTML = tableTemplate(pullResults);
    }
  };

  var loadPlayers = function() {
    lambda.invoke(indexPullParams, lambdaIndexHandler);
  };

  var lambdaCreateHandler = function(err, data) {
    if (err) {
      console.warn(err);
    } else {
      pullResults = JSON.parse(data.Payload);
      console.log('player created', pullResults);
      loadPlayers();
    }
  }

  lambda.invoke(indexPullParams, lambdaIndexHandler);

  var newPlayerForm = document.querySelector('.new_player');
  var addNewPlayer = function(ev) {
    ev.preventDefault();
    var playerName = newPlayerForm.querySelector('input[type="text"]').value;
    createPullParams['Payload'] = JSON.stringify({ name: playerName });

    lambda.invoke(createPullParams, lambdaCreateHandler);
  }

  newPlayerForm.addEventListener('submit', addNewPlayer);
  loadPlayers();
}

window.addEventListener('load', viewHandler);
