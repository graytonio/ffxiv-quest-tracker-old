class CompleteCounter extends HTMLElement {
  constructor() {
    super()
  }

  getChildQuests() {
    return Array.from(this.closest("details").querySelectorAll("input[type=checkbox]")).filter(el => !el.closest("summary"))
  }

  addListeners() {
    htmx.onLoad((e) => {
      this.updateSummary()
    })
  }

  updateSummary() {
    let quests = this.getChildQuests()

    let completedQuests = quests.filter(el => {
      return el.checked
    })

    let percentage = ((completedQuests.length/quests.length)*100).toFixed(0)

    this.textContent = `${this.getAttribute("title")} - ${completedQuests.length}/${quests.length} ${percentage}%`

    this.closest('details').querySelector("summary input[type=checkbox]").checked = (quests.length == completedQuests.length)
  }

  connectedCallback() {
    this.addListeners()
    this.updateSummary()
  }
}
customElements.define("complete-counter", CompleteCounter);
