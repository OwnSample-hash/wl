for (let i = 1; i < Number("{{len .Prods}}") + 1; i++) {
  $("#length-" + i)[0].onchange();
}
function sendPut(id) {
  $.ajax({
    method: "PUT",
    url: "/admin/products/" + id,
    headers: {
      "X-CSRF-Token": "{{.CsrfToken}}",
    },
    data: {
      name: $("#name").val(),
      description: $("#description").val(),
      one_month: $("#one_month").prop("checked") ? "on" : "off",
      life_time: $("#life_time").prop("checked") ? "on" : "off",
      price_per_month: $("#price_per_month").val(),
      price: $("#price").val(),
      discount: $("#discount").val(),
    },
    success: function () {
      location.reload();
    },
    error: function (xhr) {
      console.log(
        "Request Status: " +
          xhr.status +
          " Status Text: " +
          xhr.statusText +
          " " +
          xhr.responseText
      );
      alert("check console for error");
    },
  });
}
function deleteThis(id) {
  $.ajax({
    method: "DELETE",
    url: "/admin/products/" + id,
    headers: {
      "X-CSRF-Token": "{{.CsrfToken}}",
    },
    success: function () {
      location.reload();
    },
    error: function (xhr) {
      console.log(
        "Request Status: " +
          xhr.status +
          " Status Text: " +
          xhr.statusText +
          " " +
          xhr.responseText
      );
      alert("check console for error");
    },
  });
}
function flipThis(id) {
  $.ajax({
    method: "PATCH",
    url: "/admin/products/" + id + "/flip",
    headers: {
      "X-CSRF-Token": "{{.CsrfToken}}",
    },
    success: function () {
      location.reload();
    },
    error: function (xhr) {
      console.log(
        "Request Status: " +
          xhr.status +
          " Status Text: " +
          xhr.statusText +
          " " +
          xhr.responseText
      );
      alert("check console for error");
    },
  });
}
function editThis(
  id,
  name,
  description,
  one_month,
  life_time,
  price_per_month,
  price,
  discount
) {
  $("#submit-prod").html("Edit");
  $("#submit-prod").attr("type", "button");
  $("#name").val(name);
  $("#description").val(description);
  $("#one_month").prop("checked", one_month === "true");
  $("#life_time").prop("checked", life_time === "true");
  $("#price_per_month").val(price_per_month);
  $("#price").val(price);
  $("#discount").val(discount);
  $("#submit-prod").attr("onclick", "sendPut('" + id + "')");
}
