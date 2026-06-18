export default function calendarDays() {
  const pnlMap = {
    1: 320,
    2: -180,
    3: 0,
    4: 480,
    5: 0,
    6: 0,
    7: 0,
    8: 220,
    9: -90,
    10: 550,
    11: 180,
    12: -310,
    13: 0,
    14: 0,
    15: 90,
    16: -60,
    17: 280,
    18: 420,
    19: 0,
    20: 0,
    21: 0,
    22: 150,
    23: -240,
    24: 380,
    25: 60,
    26: 310,
    27: 0,
    28: 0,
    29: 220,
    30: -80,
  };

  const days = [];

  for (let d = 1; d <= 30; d++) {
    const pnl = pnlMap[d] ?? 0;
    let cls = "";

    if (pnl > 400) cls = "bg-success-400 text-success-900";
    else if (pnl > 100) cls = "bg-success-200 text-success-800";
    else if (pnl === 0) cls = "bg-neutral-100 text-neutral-300";
    else if (pnl > -200) cls = "bg-red-100 text-red-400";
    else cls = "bg-red-300 text-red-800";

    days.push({
      day: d,
      cls,
      label: pnl !== 0 ? (pnl > 0 ? `+$${pnl}` : `$${pnl}`) : "—",
    });
  }

  return days;
}
