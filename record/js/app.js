var teams = {
  0: {
    playerIds: [],
    score: 0,
  },
  1: {
    playerIds: [],
    score: 0,
  }
};

function bindStartGameButton() {
  var startGameButton = document.querySelector('.startGame');

  startGameButton.addEventListener('click', function() {
    [].forEach.call(document.querySelectorAll('.game'), function(gameSection) {
      gameSection.classList.add('started');
    });
    [].forEach.call(document.querySelectorAll('.pregame'), function(gameSection) {
      gameSection.classList.add('hidden');
    });
    Record.start();
  });
}

function bindEndGameButton() {
  var stopGameButton = document.querySelector('.stopGame');
  stopGameButton.addEventListener('click', function() {
    Record.stop();
  });
}

PlayersTable.init();
ScoreBoard.init();
bindStartGameButton();
bindEndGameButton();
