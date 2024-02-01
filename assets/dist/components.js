class HeaderComponent extends HTMLElement {
  constructor() {
    self = super();
    this.attachShadow({mode: "open"});
    this.shadowRoot.innerHTML = `
        <link rel="stylesheet" href="/assets/styles.css" />
        <details class="collapse my-2">
          <summary class="collapse-title text-xl bg-base-200 font-medium mb-2">Section</summary>
          <div class="collapse-content bg-base-100"><slot></slot></div>
        </details>`;
  }

  updateSummary() {
    let summary = this.shadowRoot.querySelector("summary");
    let allQuests = Array.from(this.querySelectorAll("input[type=checkbox]"));

    allQuests.forEach((el) => {
      el.addEventListener("change", () => this.updateSummary());
    });

    let completedQuests = allQuests.filter(el => el.checked);

    summary.textContent = `${this.getAttribute("title")} - ${completedQuests.length}/${allQuests.length}`;
  }

  connectedCallback() {
    this.updateSummary();
  }
}
customElements.define("collapse-section", HeaderComponent);
