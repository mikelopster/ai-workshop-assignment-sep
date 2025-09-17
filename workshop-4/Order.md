## 1. Foundation: Authentication & Wallet

* **Features:** User login, profile setup, wallet balance storage, ledger model.
* **Reason:** Every other feature depends on users having accounts, balances, and a reliable points ledger.

---

## 2. Dashboard & Transaction History

* **Features:** Wallet dashboard, recent activity, view all history, filters.
* **Reason:** Gives visibility and trust in the system. Users need to confirm their points are accurate before transfers/shopping. It’s also the baseline for QA/testing.

---

## 3. Transfer Points (P2P)

* **Features:** Select recipient, enter amount (with presets), confirm, success/failure handling.
* **Reason:** Core “value movement” feature. If this doesn’t work, QR and shopping flows have no foundation. Also critical to test balance updates in real-time.

---

## 4. QR Request / Receive Points

* **Features:** Enter amount, generate QR, share, mark as paid.
* **Reason:** Builds on transfer. Adds convenience but not critical until transfers are stable. Requires reliable balance updates already in place.

---

## 5. Shopping: Catalog → Cart → Checkout

* **Features:** Product list, add to cart, subtotal, checkout with LBK Points.
* **Reason:** Once wallet and transfers are stable, shopping can be layered on top. Uses the same ledger mechanics but with products instead of P2P. It’s more complex (cart, stock, order system) so comes after transfers.

---

## 6. Attach Customer to Order

* **Features:** Input customer name/phone at checkout.
* **Reason:** Adds metadata on top of shopping. Doesn’t block shopping itself, so can come after base order flow is working.

---

## 7. SMS Receipt

* **Features:** Send SMS after checkout (order details + confirmation).
* **Reason:** Last-mile polish for real-world use. Needs shopping orders first. Also introduces integration with external provider, so better to test late.

---

## 8. Profile / Settings Enhancements

* **Features:** Membership tier display, editable profile fields, links.
* **Reason:** Non-critical and low dependency. Can be polished at the end to improve UX.

---

# 📌 Final Sequence

1. **Authentication & Wallet (ledger)**
2. **Dashboard & History**
3. **Transfer Points**
4. **QR Request**
5. **Shopping (Catalog → Cart → Checkout)**
6. **Attach Customer**
7. **SMS Receipt**
8. **Profile/Settings polish**
