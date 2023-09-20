function myFunction() {
  // Declare variables
  var input, filter, articles, i, h2, txtValue, datalist;

  // Get the input element and its value, transformed to uppercase
  input = document.getElementById("artistSearch");
  filter = input.value.toUpperCase();

  // Get the datalist element
  datalist = document.getElementById("artistSuggestions");

  // Clear previous suggestions from datalist
  datalist.innerHTML = "";

  // Get all article elements that match your general selector
  articles = document.querySelectorAll(
    "body > main > section.s2 > div > article > a > div.front > div.musicContent > div:nth-child(1) > h2"
  );

  // Loop through all matching articles
  for (i = 0; i < articles.length; i++) {
    // Get the h2 element within the current article
    h2 = articles[i];

    // Get the text content of the h2 element
    txtValue = h2.textContent || h2.innerText;

    // Check if the article should be displayed or not
    if (txtValue.toUpperCase().indexOf(filter) > -1) {
      // Go up to the closest article element to show it
      h2.closest("article").style.display = "";

      // Add new suggestion to the datalist
      const option = document.createElement("option");
      option.value = txtValue;
      datalist.appendChild(option);
    } else {
      // Go up to the closest article element to hide it
      h2.closest("article").style.display = "none";
    }
  }
}
