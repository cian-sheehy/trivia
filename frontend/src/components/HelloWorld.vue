<template>
<v-app id="inspire">
  <div class="mt-32 container">
    <section>
      <div class="mt-32 container">
        <div v-if="pageLoading">
          <v-main>
            <v-container class="fill-height" fluid>
              <v-row align="center" justify="center">
                <v-col class="text-center">
                  <v-progress-circular :size="80" width="5" color="#e87d3d" indeterminate></v-progress-circular>
                </v-col>
              </v-row>
            </v-container>
          </v-main>
        </div>
        <div v-else>
          <v-expansion-panels v-model="panel" multiple>
            <v-expansion-panel focusable multiple v-for="(q, qIndex) in form.results" v-bind:key="q.question" class="mt-2 mb-4 rounded-lg">
              <v-expansion-panel-header>
                Q: {{q.question}}
                <template v-slot:actions v-if="q.answered">
                  <v-icon v-if="q.incorrectAnswer != ''" color="#e87d3d">mdi-alert-circle</v-icon>
                  <v-icon v-else color="teal">mdi-check</v-icon>
                </template>
              </v-expansion-panel-header>
              <v-expansion-panel-content v-if="!q.answered">
                <v-slide-item v-for="(answer) in q.incorrect_answers" v-bind:key="answer">
                  <v-btn class="mx-2 mb-3" small active-class="purple white--text" depressed rounded :disabled="q.disabled" @click="selectAnswer($event, q.correct_answer, qIndex)">
                    {{answer}}
                  </v-btn>
                </v-slide-item>
              </v-expansion-panel-content>
              <v-expansion-panel-content v-else>
                <v-btn class="mx-2 mb-3" small color="success" active-class="purple white--text" depressed rounded>
                  {{q.correct_answer}}
                </v-btn>
                <v-btn v-if="q.incorrectAnswer != ''" class="mx-2 mb-3" small color="#e87d3d" active-class="purple white--text" depressed rounded>
                  {{q.incorrectAnswer}}
                </v-btn>
              </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>
          <v-btn fab xLarge color="#e87d3d" fixed right bottom>
            {{correctAnswers +"/"+ totalAnswers}}
          </v-btn>
        </div>
      </div>
    </section>
  </div>
</v-app>
</template>

<script>
import axios from "axios";

export default {
  name: "HelloWorld",
  data() {
    return {
      loading: false,
      pageLoading: true,
      isItCorrect: null,
      panel: [],
      correctAnswers: 0,
      totalAnswers: 0,
      form: {
        question: "",
        answers: [],
        variant: [],
        selectedAnswer: ""
      }
    }
  },
  mounted: async function () {
    this.pageLoading = true
    await axios({
        method: "GET",
        url: "/api/",
        timeout: 30000,
      })
      .then((response) => {
        this.form = response.data

        for (let i = 0; i < this.form.results.length; i++) {
          this.form.results[i].incorrect_answers = this.form.results[i].incorrect_answers.sort(() => Math.random() - 0.5)
          this.$set(this.form.results[i], "answeredCount", 0)
          this.$set(this.form.results[i], "disabled", false)
          this.$set(this.form.results[i], "answered", false)
          this.$set(this.form.results[i], "incorrectAnswer", false)

          this.form.results[i].question = this.form.results[i].question.replaceAll("\u0026quot;", "\"").replaceAll("&quot;", "\"").replaceAll("&#039;", "'").replaceAll("&amp;", "&")
          this.form.results[i].correct_answer = this.form.results[i].correct_answer.replaceAll("\u0026quot;", "\"").replaceAll("&quot;", "\"").replaceAll("&#039;", "'").replaceAll("&amp;", "&")
          for (let k = 0; k < this.form.results[i].incorrect_answers.length; k++) {
            this.form.results[i].incorrect_answers[k] = this.form.results[i].incorrect_answers[k].replaceAll("\u0026quot;", "\"").replaceAll("&quot;", "\"").replaceAll("&#039;", "'").replaceAll("&amp;", "&")
          }
        }
      })
      .catch((error) => {
        console.log(error);
        this.errored = true;
      });
    this.panel = [...this.form.results].map((k, i) => i)
    this.pageLoading = false;
  },
  methods: {
    selectAnswer: function (event, correctAnswer, qIndex) {
      this.form.results[qIndex].answeredCount++
      this.totalAnswers++
      if (event.target.innerText.toLowerCase() === correctAnswer.toLowerCase()) {
        this.form.results[qIndex].answered = this.form.results[qIndex].disabled = true
        this.panel = this.panel.filter(i => i != qIndex);
        this.correctAnswers++
      } else {
        this.form.results[qIndex].answered = this.form.results[qIndex].disabled = this.form.results[qIndex].answeredCount = 1
        this.form.results[qIndex].incorrectAnswer = event.target.innerText
      }
    },
  },
  created() {
    this.loading = false
  }
}
</script>

<style scoped>
.v-btn {
  min-width: 0;
  margin-top: 1em;
}

.v-btn:hover:before,
.v-btn:focus:before {
  background-color: #21e6c1;
}
</style>
