const template = document.createElement("template");
template.innerHTML = `
<style>
  input {
    width: 100%;
    padding: 0.5rem;
    margin-top: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 0.25rem;
  }
</style>
<input type="textarea" placeholder="Type something..."/>
`;

class WYSIWYGEditor extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    const shadow = this.attachShadow({ mode: "open" });
    shadow.appendChild(template.content.cloneNode(true));
  }
}

customElements.define("wysiwyg-editor", WYSIWYGEditor);
