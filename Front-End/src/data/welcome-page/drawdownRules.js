export default [
  {
    dot: "bg-red-400",
    type: "Static Max Drawdown",
    desc: "Fixed floor from initial balance. Account can never fall below initial_balance − max_DD.",
  },
  {
    dot: "bg-gold-400",
    type: "Daily Drawdown",
    desc: "Resets at configurable timezone. Protects against single bad days wiping the challenge.",
  },
  {
    dot: "bg-primary-400",
    type: "Trailing (High-Water Mark)",
    desc: "Floor moves up as equity grows, never down. Closest to real funded account rules.",
  },
  {
    dot: "bg-success-400",
    type: "Combined Mode",
    desc: "Enforce max DD and daily DD simultaneously. One violation fails the challenge.",
  },
];
