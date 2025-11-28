const aulas = [
  { id: 1, data: "29/08/2025", titulo: "Lista de 25 exercícios de GoLang", descricao: "Entrega de uma lista com 25 exercícios de GoLang. Exercícios de nível básico para introdução à linguagem.", tipo: "exercicios" },
  { id: 2, data: "05/09/2025", titulo: "Testes de Benchmark em 5 exercícios", descricao: "Escolher 5 dos 10 primeiros exercícios para realizar testes de Benchmark. Produção de um relatório com os resultados obtidos.", tipo: "benchmark" },
  { id: 3, data: "19/09/2025", titulo: "Semana da Administração - Sem aula", descricao: "Semana da Administração. Liberados para assistir o evento.", tipo: "aula", anotacao: "Durante essa semana de aula estava ocorrendo a Semana da Administração no campus. (15/09 --> 19/09)\nFomos dispensados para assistir ao evento no dia da aula.\n\nIsso ocorreu no dia (19/09/2025)" },
  { id: 4, data: "26/09/2025", titulo: "Verificação do portfólio e atividades", descricao: "Sem conteúdo novo. Professor Alex verificou o andamento do portfólio e atividades. Presença não confirmada.", tipo: "aula", anotacao: "Nesse dia, professor Alex pediu para ver como está o andamento do portifólio até aquele momento.\nEra preciso ir até sua sala e apresentar como estava o trabalho e tirar dúvidas.\n\nIsso ocorreu no dia (26/09/2025)" },
  { id: 5, data: "03/10/2025", titulo: "Não houve aula", descricao: "Fomos liberados da aula.", tipo: "aula" },
  { id: 6, data: "10/10/2025", titulo: "Finalizar exercícios anteriores", descricao: "Solicitação para finalizar os exercícios anteriores.", tipo: "aula", anotacao: "Foi pedido para resolver alguns exercícios do Go by example durante a aula.\n\nIsso ocorreu no dia (10/10/2025)" },
  { id: 7, data: "17/10/2025", titulo: "Leitura de arquivo .ppm em GoLang", descricao: "Atividade prática em GoLang: Ler um arquivo .ppm. Exibir os valores dos pixels. Cada deve criar o seu próprio arquivo .ppm", tipo: "ppm", anotacao: "Fazer um Algoritmo que lê um arquivo .ppm (no padrão p3)" },
  { id: 8, data: "24/10/2025", titulo: "Revisão de funções e atributos", descricao: 'Aula normal. Revisão de funções e atributos usando <a href="https://gobyexample.com" target="_blank">gobyexample.com</a>. Foi pedido também para estudar diferença entre <strong>defer</strong> e <strong>panic</strong>.', tipo: "aula"},
  { id: 9, data: "31/10/2025", titulo: "Recesso - Sem aula", descricao: "Recesso conforme calendário.", tipo: "aula" },
  { id: 10, data: "07/11/2025", titulo: "Palestra de consciência negra", descricao: "Palestra de consciência negra. Teve chamada, mas não houve conteúdo da disciplina.", tipo: "aula" },
  { id: 11, data: "14/11/2025", titulo: "Formatura do antigo TADS 5 - Sem aula", descricao: "No dia houve a formatura do antigo TADS 5, fomos liberados, sem aula.", tipo: "aula" },
  { id: 12, data: "21/11/2025", titulo: "Recesso - Sem aula", descricao: "Recesso conforme calendário.", tipo: "aula" }
];

let exercicios = [];
let exerciciosBenchmark = [];
let ppmCodigo = "";
 
async function loadExercisesFromFiles() {
  const base = "anotacoes/LISTAS_EXERXICIOS/Lista1-29-08-25";
  const lista = [];
  for (let i = 1; i <= 25; i++) {
    const path = `${base}/ex${i}.go`;
    const resp = await fetch(path);
    if (!resp.ok) continue;
    const codigo = await resp.text();
    const match = codigo.match(/\/\/\s*(.+)/);
    const enunciado = match ? match[1] : `Exercício ${i}`;
    lista.push({ numero: i, enunciado, codigo });
  }
  exercicios = lista;
}

async function loadBenchmarkFiles() {
  const base = "anotacoes/LISTAS_EXERXICIOS/Lista2-05-09-25";
  const lista = [];
  for (let n = 1; n <= 6; n++) {
    const codigo = await fetch(`${base}/ex${n}_Teste/ex${n}.go`).then(r => r.text());
    const test = await fetch(`${base}/ex${n}_Teste/ex${n}_test.go`).then(r => r.text());
    lista.push({
      numero: n,
      nome: `Exercício ${n}`,
      codigo, 
      test
    });
  }
  exerciciosBenchmark = lista;
}

async function loadPPM() {
  ppmCodigo = await fetch("anotacoes/LISTAS_EXERXICIOS/Lista4-17-10-25/ArquivoPPM/read-ppm.go").then(r => r.text());
}

function renderCards() {
  const container = document.getElementById("cardsContainer");
  container.innerHTML = "";
  aulas.forEach(aula => {
    const card = document.createElement("div");
    card.className = "card";
    card.onclick = () => openCardModal(aula);
    card.innerHTML = `
      <div class="card-header">
        <h2>${aula.titulo}</h2>
        <span class="date">${aula.data}</span>
      </div>
      <div class="card-body"><p>${aula.descricao}</p></div>
      <div class="card-footer"><div class="click-hint">Clique para ver mais</div></div>
    `;
    container.appendChild(card);
  });
}

function openCardModal(aula) {
  if (aula.tipo === "exercicios") {
    openModal("exerciciosModal");
    loadExercises();
  } else if (aula.tipo === "benchmark") {
    openModal("benchmarkModal");
    loadBenchmark();
  } else if (aula.tipo === "ppm") {
    openModal("ppmModal");
    loadPPMExercise();
  } else openAulaModal(aula);
}

function openAulaModal(aula) {
  document.getElementById("aulaModalTitle").textContent = `${aula.data} - ${aula.titulo}`;
  let html = `<p><strong>${aula.descricao}</strong></p>`;
  if (aula.anotacao) {
    html += `
      <div class="code-section" style="margin-top:1.5rem;">
        <h3>Anotações</h3>
        <pre style="background:rgba(0,173,216,.1);padding:1rem;border-radius:8px;">${aula.anotacao}</pre>
      </div>`;
  }
  document.getElementById("aulaModalContent").innerHTML = html;
  openModal("aulaModal");
}

function openModal(id) {
  const m = document.getElementById(id);
  m.classList.add("active");
  m.style.display = "flex";
  document.body.style.overflow = "hidden";
}

function closeModal(id) {
  const m = document.getElementById(id);
  m.classList.remove("active");
  m.style.display = "none";
  document.body.style.overflow = "auto";
}

window.onclick = e => {
  document.querySelectorAll(".modal").forEach(m => {
    if (e.target === m) {
      m.classList.remove("active");
      m.style.display = "none";
      document.body.style.overflow = "auto";
    }
  });
};

async function loadExercises() {
  const tabs = document.getElementById("exercisesTabs");
  const content = document.getElementById("exerciseContent");

  if (!exercicios.length) await loadExercisesFromFiles();

  tabs.innerHTML = "";
  content.innerHTML = "";

  exercicios.forEach(ex => {
    const t = document.createElement("div");
    t.className = "exercise-tab";
    t.textContent = `Ex${ex.numero}`;
    t.onclick = () => showExercise(ex.numero);
    tabs.appendChild(t);
  });

  showExercise(1);
}

function showExercise(n) {
  const ex = exercicios.find(e => e.numero === n);
  if (!ex) return;

  document.querySelectorAll(".exercise-tab").forEach(t => t.classList.remove("active"));
  const tabs = document.querySelectorAll(".exercise-tab");
  if (tabs[n - 1]) tabs[n - 1].classList.add("active");

  document.getElementById("exerciseContent").innerHTML = `
    <div class="exercise-display">
      <div class="exercise-enunciado"><h3>Exercício ${ex.numero}</h3><p>${ex.enunciado}</p></div>
      <div class="exercise-code"><h3>Código:</h3><pre><code>${escapeHtml(ex.codigo)}</code></pre></div>
    </div>`;
}

async function loadBenchmark() {
  if (!exerciciosBenchmark.length) await loadBenchmarkFiles();
  const c = document.getElementById("benchmarkContent");
  c.innerHTML = "";
  exerciciosBenchmark.forEach(ex => {
    c.innerHTML += `
      <div class="benchmark-exercise">
        <h4>${ex.nome}</h4>
        <div class="benchmark-code"><pre><code>${escapeHtml(ex.codigo)}</code></pre></div>
        <div class="benchmark-test"><h5>Teste de Benchmark:</h5><pre><code>${escapeHtml(ex.test)}</code></pre></div>
      </div>`;
  });

  c.innerHTML += `
    <div class="benchmark-report">
      <h4>Relatório de Benchmark</h4>
      <p><strong>Objetivo:</strong> Realizar testes de benchmark em 6 exercícios selecionados para medir a velocidade de execução.<br><br>
      <strong>Metodologia:</strong> Os testes foram executados com múltiplas iterações automáticas pelo Go.<br><br>
      <strong>Conclusão:</strong> Os benchmarks permitem comparar o desempenho e otimizar funções Go.</p>
    </div>`;
}

async function loadPPMExercise() {
  if (!ppmCodigo) await loadPPM();
  document.getElementById("ppmCode").textContent = ppmCodigo;
}

function escapeHtml(t) {
  const d = document.createElement("div");
  d.textContent = t;
  return d.innerHTML;
}

document.addEventListener("DOMContentLoaded", renderCards);
