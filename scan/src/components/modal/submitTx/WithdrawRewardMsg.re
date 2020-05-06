[@react.component]
let make = (~setMsgsOpt) => {
  let (validator, setValidator) = React.useState(_ => EnhanceTxInput.empty);

  React.useEffect1(
    _ => {
      let msgsOpt = {
        let%Opt validatorValue = validator.value;
        Some([|TxCreator.WithdrawReward(validatorValue)|]);
      };
      setMsgsOpt(_ => msgsOpt);
      None;
    },
    [|validator|],
  );

  <>
    <EnhanceTxInput
      width=360
      inputData=validator
      setInputData=setValidator
      parse=Address.fromBech32Opt
      msg="Undelegate from"
      errMsg="Invalid Address"
    />
  </>;
};
