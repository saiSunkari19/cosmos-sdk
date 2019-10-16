# Auth and Bank Redesign Proposal

1. Bank should have it's own store that is a mapping of `Address->sdk.Coins` (really in the future, this should be a more generic `sdk.Assets` but that's fine for now.).  `sdk.BaseAccount` should not have a `balance`, all balances are kept in the `BankKeeper`.  The reason this was not done originally is because of the worries that this will require an addition state read during the AnteHandler in which we have to pull the `sdk.BaseAccount` for signature and sequence checking, and then the balance for fee deduction.  While this is true, I think the modularity and simplicity that balances in the BankKeeper make the single state read addition worth it.
2. The Supply tracking 