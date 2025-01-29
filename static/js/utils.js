function showPrice(id, type, pricePerMonth, price, discount, discountPrice, discountPerMonth) {
  if (Number(discount) > 0) {
    if (type === 'one_month') {
      $("#price-discount-" + id).html(pricePerMonth);
      $("#price-discounted-" + id).html(discountPerMonth);
      $("#price-modal-discount-" + id).html(pricePerMonth);
      $("#price-modal-discounted-" + id).html(discountPerMonth);
    } else {
      $("#price-discount-" + id).html(price);
      $("#price-discounted-" + id).html(discountPrice);
      $("#price-modal-discount-" + id).html(price);
      $("#price-modal-discounted-" + id).html(discountPrice);
    }
  } else {
    if (type === 'one_month') {
      $("#price-" + id).html(pricePerMonth);
      $("#price-modal-" + id).html(pricePerMonth);
    } else {
      $("#price-" + id).html(price);
      $("#price-modal-" + id).html(price);
    }
  }
}
