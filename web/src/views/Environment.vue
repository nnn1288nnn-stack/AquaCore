<template>
  <div>
    <h1>🌊 環境監測</h1>
    <p>這裡將顯示養殖環境的即時數據和歷史趨勢。</p>

    <div class="row">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5>水質參數</h5>
          </div>
          <div class="card-body">
            <p>即時監測水溫、鹽度、pH 值等參數</p>
          </div>
        </div>
      </div>

      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5>趨勢分析</h5>
          </div>
          <div class="card-body">
            <p>查看歷史數據趨勢和預測</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Environment'
}
</script>