ScoreBoard = function() {
  var buttons = document.querySelectorAll('.game table .goals .btn');

  var handleScoreChange = function(button) {
    var valueButton = button.parentNode.querySelector('.value');
    var teamId = parseInt(valueButton.getAttribute('data-team-id'), 10);
    var changeBy = 0;
    if (button.classList.contains('minus') && teams[teamId].score > 0) {
      changeBy -= 1;
    } else if (button.classList.contains('plus')) {
      changeBy += 1;
    }
    teams[teamId].score += changeBy;
    valueButton.innerHTML = teams[teamId].score;
  };

  var init = function() {
    [].forEach.call(buttons, function(button) {
      button.addEventListener('click', function() { handleScoreChange(button) });
    });
  }

  return {
    init: init
  }
}();
