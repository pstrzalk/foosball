var PlayersTable = function () {
  var table = document.querySelector('.players');

  var players = [
    { id: 1, name: 'Joe Black' },
    { id: 2, name: 'Peter Green' },
    { id: 3, name: 'John Zeo' },
    { id: 4, name: 'Xavier Keller' },
    { id: 5, name: 'Bo Bugger' },
    { id: 6, name: 'Chloe Dowe' },
  ];

  var bindPlayerTable = function() {
    playerTableCells = PlayersTable.table.querySelectorAll('tr[data-id] td');
    [].forEach.call(playerTableCells, function(playerTableCell) {
      playerTableCell.removeEventListener('click', setPlayerTeam);
      playerTableCell.addEventListener('click', setPlayerTeam);
    });
  }

  var setPlayerTeam = function(ev) {
    var playerTableCell = ev.currentTarget;
    var playerRow = playerTableCell.parentNode;
    var playerId = parseInt(playerRow.getAttribute('data-id'), 10);
    var teamId = parseInt(playerTableCell.getAttribute('data-team-id'), 10);
    var startGameButton = document.querySelector('.startGame');

    teams[0].playerIds = teams[0].playerIds.filter(function(id) { return id != playerId });
    teams[1].playerIds = teams[1].playerIds.filter(function(id) { return id != playerId });
    playerRow.querySelector('[data-team-id="0"]').innerHTML = '';
    playerRow.querySelector('[data-team-id="1"]').innerHTML = '';

    if ((teamId == 0 || teamId == 1) && teams[teamId].playerIds.length < 2) {
      teams[teamId].playerIds.push(playerId);
      playerRow.querySelector(`[data-team-id="${teamId}"]`).innerHTML = '<i class="nes-logo">&nbsp;</i>'
    }

    startGameButton.disabled = teams[0].playerIds.length < 2 || teams[1].playerIds.length < 2;
  }

  var init = function() {
    var odd = false;
    players.forEach(function(player) {
      table.innerHTML += `
        <tr data-id="${player.id}" class="playerRow ${odd ? 'odd' : ''}">
          <td data-team-id="0" class="team"></td>
          <td class="js-off">${player.name}</td>
          <td data-team-id="1" class="team"></td>
        </tr>`;
      odd = !odd;
    });
    bindPlayerTable();
  }

  return {
    init: init,
    table: table,
  };
}();
