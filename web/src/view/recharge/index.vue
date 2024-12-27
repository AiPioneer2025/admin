  <template>

    <h1>余额充值</h1>
    <el-form :model="form" :rules="rules" ref="rechargeForm" label-width="100px" @keyup.enter="handleRecharge">
      <el-form-item class="mb-6" label="充值金额" prop="rechargeAmount">
        <el-input v-model="form.rechargeAmount" placeholder="请输入充值金额"></el-input>
      </el-form-item>
      <el-form-item class="mb-6">
        <el-button 
          class="shadow shadow-active h-11 w-full" 
          type="primary" 
          @click="handleRecharge"
          >确认充值</el-button>
      </el-form-item>
    </el-form>
  </template>
  
  <script>
  export default {
    data() {
      return {
        form: {
          rechargeAmount: ''
        },
        rules: {
          rechargeAmount: [
            { required: true, message: '请输入充值金额', trigger: 'blur' },
            { min: 1, message: '充值金额必须大于0', trigger: 'blur' }
          ]
        }
      };
    },
    methods: {
      handleRecharge() {
        const txamt = this.rechargeAmount * 100; // 将金额转换为分
        const out_trade_no = `recharge_${new Date().getTime()}`; // 生成唯一的外部交易号
        const txdtm = new Date().toISOString(); // 下单时间
  
        // 构建请求参数
        const params = {
          appcode: 'YOUR_APP_CODE', // 替换为 QFPay 提供的 API 凭证
          sign_type: 'sha256',
          sign: 'GENERATED_SIGNATURE', // 替换为生成的签名
          paysource: 'remotepay_checkout',
          txamt: txamt,
          txcurrcd: 'CNY', // 假设货币代码为人民币
          out_trade_no: out_trade_no,
          txdtm: txdtm,
          return_url: 'https://yourdomain.com/return/success', // 替换为支付成功后的重定向URL
          failed_url: 'https://yourdomain.com/return/failed', // 替换为支付失败后的重定向URL
          notify_url: 'https://yourdomain.com/notify/success', // 替换为异步通知URL
          goods_name: '余额充值',
          lang: 'zh-cn',
        };
  
        // 将参数转换为查询字符串并发起请求
        const queryString = new URLSearchParams(params).toString();
        window.location.href = `https://test-openapi-hk.qfapi.com/checkstand/#/?${queryString}`;
      }
    }
  };
  </script>
  
  <style>
  .balance-recharge {
    /* 样式根据需要添加 */
  }
  </style>
  