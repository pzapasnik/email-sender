const template = document.createElement("template");
template.innerHTML = `
<style>
  .container {
    height: 100%;
    display: flex;
    flex-direction: column;
    padding: 2px;
  }
  .toolbar {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    padding: 2px;
    border: 1px solid black;
  }
  .button {
    padding: 2px;
    margin: 2px;
    border: 1px solid black;
    border-radius: 4px;
  }
  /* .message { */
  /*   height: 100%; */
  /* } */
  /* :host { */
  /*   display: block; */
  /*   height: 100%; */
  /* } */
</style>
<div class="container">
  <slot name="label"></slot>
  <slot name="toolbar"></slot>
  <slot id="wysiwyg-editor-message"class="message" name="message"></slot>
</div>
`;

class WYSIWYGEditor extends HTMLElement {
  constructor() {
    super();
    const shadow = this.attachShadow({ mode: "open" });
    shadow.appendChild(template.content.cloneNode(true));
  }
}

customElements.define("wysiwyg-editor", WYSIWYGEditor);
