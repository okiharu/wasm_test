var clickCount = 0;
var avgTime = 0;
var totalTime = 0;

function findByJs() {
  //start
  const start = performance.now();

  var str = document.getElementById('randomStr').value;
  var counter = 0;
  for (var i = 0; i < str.length - 5; i++) {
    if (str.substring(i, i + 6) == 'SUPINF') {
      counter++;
    }
  }
  document.getElementById('result').value = counter;
  //end
  const end = performance.now();
  var time = end - start;
  document.getElementById('spdJs').value = time;

  culcAvg(time);
}

function culcAvg(time) {
  clickCount++;
  totalTime += time;
  document.getElementById('avgJs').value = totalTime / clickCount;
}

function init() {
  clickCount = 0;
  avgTime = 0;
  totalTime = 0;
}
