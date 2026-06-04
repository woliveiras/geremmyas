(function () {
  const listBtn = document.querySelector('[data-view-btn="list"]');
  const boardBtn = document.querySelector('[data-view-btn="board"]');
  const listPanel = document.querySelector('[data-view-panel="list"]');
  const boardPanel = document.querySelector('[data-view-panel="board"]');
  if (listBtn && boardBtn && listPanel && boardPanel) {
    listBtn.addEventListener('click', () => {
      listPanel.hidden = false;
      boardPanel.hidden = true;
      listBtn.classList.add('active');
      boardBtn.classList.remove('active');
    });
    boardBtn.addEventListener('click', () => {
      listPanel.hidden = true;
      boardPanel.hidden = false;
      boardBtn.classList.add('active');
      listBtn.classList.remove('active');
    });
  }
  const filter = document.getElementById('phase-filter');
  if (filter) {
    filter.addEventListener('change', () => {
      const phase = filter.value;
      let any = false;
      document.querySelectorAll('.board-card').forEach((card) => {
        const show = !phase || card.dataset.phase === phase;
        card.style.display = show ? '' : 'none';
        if (show) any = true;
      });
      const msg = document.getElementById('board-no-matches');
      if (msg) msg.hidden = any;
    });
  }
  const depToggle = document.getElementById('show-deprecated');
  const boardCols = document.querySelector('.board-columns');
  if (depToggle && boardCols) {
    depToggle.addEventListener('change', () => {
      boardCols.classList.toggle('show-deprecated', depToggle.checked);
    });
  }
})();
