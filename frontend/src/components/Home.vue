<template>
<div>
  <div class="mt-32 container">
    <section>
      <div class="mt-32 container">
        <b-form>
          <div v-if="pageLoading">
            <b-spinner variant="primary" style="width: 3rem; height: 3rem;" label="Spinning"></b-spinner>
          </div>
          <div v-else>
            <b-card v-for="(q, qIndex) in form.results" v-bind:key="q.question" class="mt-3">
              <b-card-header>
                {{q.question}}
                <b-badge variant="primary" pill>{{2 - q.answeredCount}}</b-badge>
              </b-card-header>
              <b-list-group v-for="(answer) in q.incorrect_answers" v-bind:key="answer">
                <b-list-group-item button :disabled="q.disabled" :variant="(answer === q.correct_answer) && (q.incorrect_answers.length === 1) ? 'success' : 'null'" @click="selectAnswer($event, q.correct_answer, qIndex)">{{answer}}</b-list-group-item>
              </b-list-group>
              <template v-if="q.showFooter" variant="success" v-slot:footer>
                <small v-if="q.disabled" class="text-muted">
                  <b-icon icon="x-circle" scale="2" variant="danger"></b-icon>
                  The correct answer is {{q.correct_answer}}
                </small>
                <small v-else class="text-muted">
                  <b-icon icon="patch-check" scale="3" variant="success"></b-icon>
                  Awesome job, you've guessed correctly on the first attempt
                </small>
              </template>
            </b-card>
          </div>
        </b-form>
      </div>
    </section>
  </div>
</div>
</template>

<script>
import axios from "axios";

export default {
  name: "home",
  data() {
    return {
      loading: false,
      pageLoading: true,
      isItCorrect: null,
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
          this.$set(this.form.results[i], "showFooter", false)
        }
      })
      .catch((error) => {
        console.log(error);
        this.errored = true;
      });
    this.pageLoading = false;
  },
  methods: {
    selectAnswer: function (event, correctAnswer, qIndex) {
      if (event.target.innerText === correctAnswer) {
        this.form.results[qIndex].incorrect_answers = []
        this.form.results[qIndex].incorrect_answers.push(correctAnswer)
        this.form.results[qIndex].showFooter = this.form.results[qIndex].answeredCount === 0
      } else {
        this.form.results[qIndex].answeredCount++
        this.form.results[qIndex].showFooter = this.form.results[qIndex].disabled = this.form.results[qIndex].answeredCount >= 2
      }
    }
  },
  created() {
    this.loading = false
  }
}
</script>

<style>
h3 {
  margin-bottom: 5%;
}

@media screen and (min-width: 700px) {
  .b-card {
    max-width: 20rem;
  }
}

@media screen and (max-width: 699px) {
  .b-card {
    width: 100%;
  }
}
</style>
