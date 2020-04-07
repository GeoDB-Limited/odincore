type block = {
  id: int,
  proposer: string,
};

module Styles = {
  open Css;

  let block = (i, ID.Block.ID(bh)) =>
    style([
      position(`absolute),
      backgroundColor(white),
      padding4(~top=`px(14), ~left=`px(10), ~right=`px(18), ~bottom=`px(16)),
      marginBottom(`px(3)),
      boxShadow(Shadow.box(~x=`zero, ~y=`px(2), ~blur=`px(2), Css.rgba(0, 0, 0, 0.05))),
      width(`px(120)),
      cursor(i == 0 || i == 11 ? auto : `pointer),
      opacity(i == 0 || i == 11 ? 0. : 1.),
      transform(
        `translate((
          `px(bh mod 2 == 1 ? 125 : 0),
          `px(i mod 2 == 1 ? i / 2 * 85 : (i + 1) / 2 * 85 - 42),
        )),
      ),
      transition(~duration=800, "all"),
    ]);

  let rightCol = style([marginLeft(`px(-3))]);

  let topicBar =
    style([
      width(`percent(100.)),
      display(`flex),
      flexDirection(`row),
      justifyContent(`spaceBetween),
    ]);

  let vFlex =
    style([display(`flex), flexDirection(`column), minWidth(`px(245)), position(`relative)]);

  let seeAll = style([display(`flex), flexDirection(`row), cursor(`pointer)]);
  let cFlex = style([display(`flex), flexDirection(`column)]);
  let amount =
    style([fontSize(`px(20)), lineHeight(`px(24)), color(Colors.gray8), fontWeight(`bold)]);
  let rightArrow = style([width(`px(25)), marginTop(`px(17)), marginLeft(`px(16))]);
};

let renderBlock = (i: int, b: BlockSub.t) =>
  <div
    key={b.height |> ID.Block.toString}
    className={Styles.block(i, b.height)}
    onClick={_ => Route.redirect(b.height |> ID.Block.getRoute)}>
    <TypeID.Block id={b.height} />
    <VSpacing size=Spacing.md />
    <Text value="PROPOSED BY" block=true size=Text.Xs color=Colors.gray7 spacing={Text.Em(0.1)} />
    <VSpacing size={`px(1)} />
    <Text
      block=true
      value={b.validator.moniker}
      weight=Text.Bold
      ellipsis=true
      height={Text.Px(15)}
      spacing={Text.Em(0.02)}
    />
  </div>;

let getDummyBlock = blocksCount =>
  BlockSub.{
    height: ID.Block.ID(blocksCount + 1),
    hash: "" |> Hash.fromHex,
    validator:
      ValidatorSub.Mini.{
        consensusAddress: "",
        operatorAddress: "" |> Address.fromHex,
        moniker: "",
      },
    timestamp: 0. |> MomentRe.momentWithTimestampMS,
    txn: 0,
  };

[@react.component]
let make = () =>
  {
    let blocksSub = BlockSub.getList(~pageSize=11, ~page=1, ());

    let%Sub realBlocks = blocksSub;
    let blocksCount =
      realBlocks
      ->Belt_Array.get(0)
      ->Belt_Option.map(({height: ID.Block.ID(x)}) => x)
      ->Belt_Option.getExn;

    let blocksWithDummy = realBlocks |> Belt_Array.concat([|getDummyBlock(blocksCount)|]);

    <>
      <div className=Styles.topicBar>
        <Text value="Latest Blocks" size=Text.Xxl weight=Text.Bold block=true color=Colors.gray8 />
      </div>
      <VSpacing size=Spacing.lg />
      <VSpacing size=Spacing.sm />
      <div className=Styles.seeAll onClick={_ => Route.redirect(Route.BlockHomePage)}>
        <div className=Styles.cFlex>
          <span className=Styles.amount> {blocksCount |> Format.iPretty |> React.string} </span>
          <VSpacing size=Spacing.xs />
          <Text
            value="ALL BLOCKS"
            size=Text.Sm
            color=Colors.bandBlue
            spacing={Text.Em(0.05)}
            weight=Text.Medium
          />
        </div>
        <img src=Images.rightArrow className=Styles.rightArrow />
      </div>
      <VSpacing size=Spacing.lg />
      <Row alignItems=`initial>
        <div className=Styles.vFlex>
          {blocksWithDummy
           ->Belt_Array.mapWithIndex((i, block) => renderBlock(i, block))
           ->React.array}
        </div>
      </Row>
    </>
    |> Sub.resolve;
  }
  |> Sub.default(_, React.null);
