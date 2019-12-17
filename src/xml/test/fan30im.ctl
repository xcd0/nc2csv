<?xml version="1.0"?>
<VcControl Version="8.2" Unit="inch">
  <Information>
    <InfoUserName>CGTech</InfoUserName>
    <InfoDateTime>January 03, 2019</InfoDateTime>
    <InfoVericutVersion>8.2.1.2</InfoVericutVersion>
  </Information>
  <COMMENT>Library File: fan30im.ctl, Revision 1.1, Issued: 03/01/2019 09:38:18</COMMENT>
  <ControlInitFile File="fan30im.ini" />
  <ConstSurfSpeed Rapid="off" MaxRPM="1000" />
  <Sync Method="none">
  </Sync>
  <Words>
    <Word Name="#" Type="variable_tag">
    </Word>
    <Word Name="$" Type="skip">
    </Word>
    <Word Name="$$" Type="skip">
    </Word>
    <Word Name="%" Type="skip">
    </Word>
    <Word Name="(" Type="begin_comment">
    </Word>
    <Word Name=")" Type="end_comment">
    </Word>
    <Word Name="*" Type="multiplication">
    </Word>
    <Word Name="+" Type="addition">
    </Word>
    <Word Name="," Type="separator">
    </Word>
    <Word Name=",C" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="4.3" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name=",R" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="4.3" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="-" Type="subtraction">
    </Word>
    <Word Name="/" Type="conditional">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
      <Function Name="AtanCondDivide"/>
    </Word>
    <Word Name=":" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="4.0" MetricType="decimal" MetricFormat="4.0" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="&lt;" Type="macro_call">
      <Value Type="alpha_numeric"/>
    </Word>
    <Word Name="&lt;&lt;" Type="concatenate">
    </Word>
    <Word Name="=" Type="assignment">
    </Word>
    <Word Name="&gt;" Type="macro_call">
      <Value Type="none"/>
    </Word>
    <Word Name="[" Type="left_precedence">
    </Word>
    <Word Name="\N" Type="end_block">
    </Word>
    <Word Name="]" Type="right_precedence">
    </Word>
    <Word Name="^" Type="power">
    </Word>
    <Word Name="A" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="3.4" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="ABS" Type="function_call">
      <Function Name="abs" Type="normal"/>
    </Word>
    <Word Name="ACOS" Type="function_call">
      <Function Name="acos_d" Type="normal"/>
    </Word>
    <Word Name="AND" Type="bitwise_and">
    </Word>
    <Word Name="ASIN" Type="function_call">
      <Function Name="asin_d" Type="normal"/>
    </Word>
    <Word Name="ATAN" Type="function_call">
      <Function Name="atan2_d" Type="normal"/>
    </Word>
    <Word Name="ATN" Type="function_call">
      <Function Name="atan2_d" Type="normal"/>
    </Word>
    <Word Name="B" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="3.4" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="BLANK" Type="macro_call">
      <Value Type="alpha_numeric"/>
    </Word>
    <Word Name="C" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="3.4" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="CLOSE" Type="macro_call">
      <Value Type="none"/>
    </Word>
    <Word Name="COS" Type="function_call">
      <Function Name="cos_d" Type="normal"/>
    </Word>
    <Word Name="D" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.0" MetricType="decimal" MetricFormat="3.0" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="DEFINE" Type="S840D_define">
    </Word>
    <Word Name="DO" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="DOORS" Type="macro_call">
      <Value Type="none"/>
    </Word>
    <Word Name="DPRNT" Type="macro_call">
      <Value Type="text_string"/>
    </Word>
    <Word Name="E" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="END" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="EQ" Type="equal">
    </Word>
    <Word Name="EXP" Type="function_call">
      <Function Name="exp" Type="normal"/>
    </Word>
    <Word Name="F" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="4.5" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="FIX" Type="function_call">
      <Function Name="IntTruncAdj" Type="normal"/>
    </Word>
    <Word Name="FUP" Type="function_call">
      <Function Name="AbsCeil" Type="normal"/>
    </Word>
    <Word Name="G" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="2.1" MetricType="decimal" MetricFormat="2.1" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="GE" Type="greater_equal">
    </Word>
    <Word Name="GOTO" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="GT" Type="greater_than">
    </Word>
    <Word Name="H" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.3" MetricType="decimal" MetricFormat="3.3" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="I" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.8" MetricType="decimal" MetricFormat="4.8" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="IF" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="J" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.8" MetricType="decimal" MetricFormat="4.8" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="K" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.8" MetricType="decimal" MetricFormat="4.8" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="L" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="4.0" MetricType="decimal" MetricFormat="4.0" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="LE" Type="less_equal">
    </Word>
    <Word Name="LENGTH" Type="function_call">
      <Function Name="STRLEN" Type="normal"/>
    </Word>
    <Word Name="LN" Type="function_call">
      <Function Name="ln" Type="normal"/>
    </Word>
    <Word Name="LT" Type="less_than">
    </Word>
    <Word Name="M" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.0" MetricType="decimal" MetricFormat="3.0" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="MP03403" Type="variable_name">
    </Word>
    <Word Name="MP_03403" Type="variable_name">
    </Word>
    <Word Name="N" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="4.0" MetricType="decimal" MetricFormat="4.0" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="NE" Type="not_equal">
    </Word>
    <Word Name="NUMBER" Type="function_call">
      <Function Name="NUMBER" Type="normal"/>
    </Word>
    <Word Name="O" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="4.0" MetricType="decimal" MetricFormat="4.0" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="OPEN" Type="macro_call">
      <Value Type="none"/>
    </Word>
    <Word Name="OR" Type="bitwise_or">
    </Word>
    <Word Name="P" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="4.4" MetricType="decimal" MetricFormat="4.3" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="PCLOS" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="POPEN" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="Q" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="4.3" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="R" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="4.3" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="RND" Type="function_call">
      <Function Name="Round" Type="normal"/>
    </Word>
    <Word Name="ROUND" Type="function_call">
      <Function Name="Round" Type="normal"/>
    </Word>
    <Word Name="S" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="SETVN" Type="Fan_setvn">
    </Word>
    <Word Name="SIN" Type="function_call">
      <Function Name="sin_d" Type="normal"/>
    </Word>
    <Word Name="SQRT" Type="function_call">
      <Function Name="sqrt" Type="normal"/>
    </Word>
    <Word Name="SUBSTR" Type="function_call">
      <Function Name="SUBSTR" Type="normal"/>
    </Word>
    <Word Name="T" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="2.0" MetricType="decimal" MetricFormat="2.0" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="TAN" Type="function_call">
      <Function Name="tan_d" Type="normal"/>
    </Word>
    <Word Name="THEN" Type="ignore">
    </Word>
    <Word Name="U" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="UNBLANK" Type="macro_call">
      <Value Type="alpha_numeric"/>
    </Word>
    <Word Name="V" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="W" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="WHILE" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="" MetricType="decimal" MetricFormat="" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="X" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="5.4" Multiply="on" Multiplier="1"/>
    </Word>
    <Word Name="XOR" Type="bitwise_xor">
    </Word>
    <Word Name="Y" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="5.4" Multiply="off" Multiplier="1"/>
    </Word>
    <Word Name="Z" Type="macro_call">
      <Value Type="numeric" InchType="decimal" InchFormat="3.4" MetricType="decimal" MetricFormat="5.4" Multiply="off" Multiplier="1"/>
    </Word>
  </Words>
  <NCSubstitutions>
  </NCSubstitutions>
  <OPSubstitutions>
  </OPSubstitutions>
  <SuperGroup Name="Specials">
    <Group Name="G" Type="word" Description="Manual Guide i Probing Cycles" Range="2020-2040" Field="0">
      <ConditionMacro>
        <Variable Name="4010" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4210" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="SubroutineName" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="MacroVarState2" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="CallNCMacro" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name=":" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="SubroutineSequence" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="N" Type="word" Description="Sequence number" Field="0">
      <ConditionMacro>
        <Variable Name="4114" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4314" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="Sequence" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="&lt;" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="M" Range="98"/>
        <Macro Name="SubroutineNameText" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="M" Range="198"/>
        <Macro Name="CallSubName" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="65 66 66.1"/>
        <Macro Name="SubroutineNameText" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="SubroutineSequence" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="&gt;" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="SubroutineSequence" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="O" Type="word" Description="Subroutine Sequence number" Field="0">
      <ConditionMacro>
        <Variable Name="4115" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4315" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="SubroutineSequenceEnd" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
  </SuperGroup>
  <SuperGroup Name="Branching">
    <Group Name="DO" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="DoLoop" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="END" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="EndLoop" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="GOTO" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="GotoJumpForwardBackwardStart" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="IF" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="IfCheck" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="WHILE" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="WhileLoop" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
  </SuperGroup>
  <SuperGroup Name="States">
    <Group Name="G" Type="word" Description="NC macro call" Range="65" Field="0">
      <ConditionMacro>
        <Macro Name="MacroVarState" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="CallNCMacro" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="NC macro call continuous with motion" Range="66" Field="0">
      <ConditionMacro>
        <Variable Name="4012" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4212" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MacroVarState2" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="MacroVarState" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="CallNCMacroMotion" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="NC macro call continuous per block" Range="66.1" Field="0">
      <ConditionMacro>
        <Variable Name="4012" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4212" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MacroVarState2" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="MacroVarState" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="CallNCMacroBlock" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="NC macro call continuous cancel" Range="67" Field="0">
      <ConditionMacro>
        <Variable Name="4012" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4212" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MacroVarState2" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="MacroVarState" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="CallNCMacroCancel" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Spindle on CW" Range="3" Field="0">
      <ConditionMacro>
        <Macro Name="ActiveSpindleActiveTool" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ActiveSpindleDir" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="CW"/>
        </Macro>
        <Macro Name="ActiveSpindleOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Spindle on CCW" Range="4" Field="0">
      <ConditionMacro>
        <Macro Name="ActiveSpindleActiveTool" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ActiveSpindleDir" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="CCW"/>
        </Macro>
        <Macro Name="ActiveSpindleOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Spindle off" Range="5" Field="0">
      <ConditionMacro>
        <Macro Name="ActiveSpindleCompName" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Spindle"/>
        </Macro>
        <Macro Name="ActiveSpindleOnOff" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Rapid positioning" Range="0" Field="0">
      <ConditionMacro>
        <Variable Name="4001" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4201" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MotionRapid" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Linear interpolation" Range="1" Field="0">
      <ConditionMacro>
        <Variable Name="4001" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4201" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MotionLinear" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Circular/helical interpolation CW" Range="2" Field="0">
      <ConditionMacro>
        <CondWord Logical="and_not" Name="L"/>
        <CondWord Logical="and_not" Name="Q"/>
        <Variable Name="4001" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4201" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MotionCW" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4001" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4201" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MotionCWSpiral" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Circular/helical interpolation CCW" Range="3" Field="0">
      <ConditionMacro>
        <CondWord Logical="and_not" Name="L"/>
        <CondWord Logical="and_not" Name="Q"/>
        <Variable Name="4001" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4201" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MotionCCW" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4001" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4201" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MotionCCWSpiral" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="High speed machining" Range="5" Field="0">
      <ConditionMacro>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Multibuffer" Range="5.1" Field="0">
      <ConditionMacro>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="NURBS interpolation" Range="6.2" Field="0">
      <ConditionMacro>
        <Variable Name="4001" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4201" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MotionNurbs" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Cylindrical Interpolation" Range="7.1" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="A" Range="0"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="A"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="B" Range="0"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="B"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="C" Range="0"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="C"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="Unsupported" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Exact stop check" Range="9" Field="0">
      <ConditionMacro>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Polar interpolation on" Range="12.1 112" Field="0">
      <ConditionMacro>
        <Variable Name="4021" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4221" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="WorkCoordCValue" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#5026"  Units="0"/>
        </Macro>
        <Macro Name="SiemensShiftOffsetB" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="PolarInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Polar interpolation off" Range="13.1 113" Field="0">
      <ConditionMacro>
        <Variable Name="4021" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4221" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="WorkCoordCValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="SiemensShiftOffsetB" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="PolarInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Polar coordinates command cancel" Range="15" Field="0">
      <ConditionMacro>
        <Variable Name="4017" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4217" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="SetPolarInput" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Polar coordinates command" Range="16" Field="0">
      <ConditionMacro>
        <Variable Name="4017" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4217" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="SetPolarInput" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="XY plane select" Range="17" Field="0">
      <ConditionMacro>
        <Variable Name="4002" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4202" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MotionPlaneXY" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="ZX plane select" Range="18" Field="0">
      <ConditionMacro>
        <Variable Name="4002" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4202" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MotionPlaneZX" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="YZ plane select" Range="19" Field="0">
      <ConditionMacro>
        <Variable Name="4002" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4202" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="MotionPlaneYZ" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Inch input" Range="20" Field="0">
      <ConditionMacro>
        <Variable Name="4006" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4206" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="UnitsInch" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Metric input" Range="21" Field="0">
      <ConditionMacro>
        <Variable Name="4006" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4206" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="UnitsMetric" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Stored stroke check function on" Range="22" Field="0">
      <ConditionMacro>
        <Variable Name="4006" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4206" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Stored stroke check function off" Range="23" Field="0">
      <ConditionMacro>
        <Variable Name="4006" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4206" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Spindle speed fluctuation detection/ off" Range="25 26" Field="0">
      <ConditionMacro>
        <Variable Name="4019" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4219" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Probe" Range="31" Field="0">
      <ConditionMacro>
        <Macro Name="Probe" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="RETURN=PROBE_TOUCHED"/>
        </Macro>
        <Macro Name="BlockFinish" Scan="no" AfterMotion="yes">
        </Macro>
        <Variable Name="5061" Scan="no" AfterMotion="yes">
          <Override  Type="expression"  Value="#5041"  Units="0"/>
        </Variable>
        <Variable Name="5062" Scan="no" AfterMotion="yes">
          <Override  Type="expression"  Value="#5042"  Units="0"/>
        </Variable>
        <Variable Name="5063" Scan="no" AfterMotion="yes">
          <Override  Type="expression"  Value="#5043"  Units="0"/>
        </Variable>
        <Variable Name="5064" Scan="no" AfterMotion="yes">
          <Override  Type="expression"  Value="#5044"  Units="0"/>
        </Variable>
        <Variable Name="5065" Scan="no" AfterMotion="yes">
          <Override  Type="expression"  Value="#5045"  Units="0"/>
        </Variable>
        <Variable Name="5066" Scan="no" AfterMotion="yes">
          <Override  Type="expression"  Value="#5046"  Units="0"/>
        </Variable>
        <Variable Name="5067" Scan="no" AfterMotion="yes">
          <Override  Type="expression"  Value="#5047"  Units="0"/>
        </Variable>
        <Variable Name="5068" Scan="no" AfterMotion="yes">
          <Override  Type="expression"  Value="#5048"  Units="0"/>
        </Variable>
        <Variable Name="5069" Scan="no" AfterMotion="yes">
          <Override  Type="expression"  Value="#5049"  Units="0"/>
        </Variable>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Normal direction control cancel mode" Range="40.1" Field="0">
      <ConditionMacro>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tool length compensation positive " Range="43" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="0"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="1"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tool length compensation in tool axis direction (5-axis)" Range="43.1" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="0"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
        <Macro Name="RtcpOn" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="1"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpOn" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpOn" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tool tip/center point control" Range="43.4" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="0"/>
        <CondVariable Logical="and" Name="MP_19680" Range="2"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="1"/>
        <CondVariable Logical="and" Name="MP_19680" Range="2"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="0"/>
        <CondVariable Logical="and" Name="MP_19680" Range="12 21"/>
        <CondVariable Logical="and" Name="MP_19696[5]" Range="0"/>
        <Macro Name="RpcpInput" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="0"/>
        <CondVariable Logical="and" Name="MP_19680" Range="12 21"/>
        <CondVariable Logical="and" Name="MP_19696[5]" Range="1"/>
        <Macro Name="RpcpInput" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="1"/>
        <CondVariable Logical="and" Name="MP_19680" Range="12 21"/>
        <CondVariable Logical="and" Name="MP_19696[5]" Range="0"/>
        <Macro Name="RpcpInput" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="1"/>
        <CondVariable Logical="and" Name="MP_19680" Range="12 21"/>
        <CondVariable Logical="and" Name="MP_19696[5]" Range="1"/>
        <Macro Name="RpcpInput" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Vector Programming" Range="43.5" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="0"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="1"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RpcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Cutting point control" Range="43.8" Field="0">
      <ConditionMacro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="RtcpOn" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RtcpMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="UpdateRotaryOffsets" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tool length compensation negative " Range="44" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="0"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="1"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tool offset increase" Range="45" Field="0">
      <ConditionMacro>
        <Macro Name="CutterCompLong" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tool offset decrease" Range="46" Field="0">
      <ConditionMacro>
        <Macro Name="CutterCompShort" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tool offset double increase" Range="47" Field="0">
      <ConditionMacro>
        <Macro Name="CutterCompDoubleLong" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tool offset double decrease" Range="48" Field="0">
      <ConditionMacro>
        <Macro Name="CutterCompDoubleShort" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tool length compensation cancel / off" Range="49" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="28"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="Cancel3dToolOffset" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ShortLongOffsetReset" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="0"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="Cancel3dToolOffset" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ShortLongOffsetReset" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05006[6]" Range="1"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="Cancel3dToolOffset" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ShortLongOffsetReset" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="Cancel3dToolOffset" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ShortLongOffsetReset" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGageOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
          <Override  Type="text"  Value="IMMEDIATE"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Scaling cancel" Range="50" Field="0">
      <ConditionMacro>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Polygon turning cancel" Range="50.2" Field="0">
      <ConditionMacro>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Mirror cancel" Range="50.1" Field="0">
      <ConditionMacro>
        <Variable Name="4022" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4222" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description=" Mirror" Range="51.1" Field="0">
      <ConditionMacro>
        <Variable Name="4022" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4222" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Machine coordinate system" Range="53" Field="0">
      <ConditionMacro>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tool axis direction control" Range="53.1" Field="0">
      <ConditionMacro>
        <Macro Name="WPMoveOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WPMatrix" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="6"  Units="0"/>
        </Macro>
        <Macro Name="WPCreate" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
          <Override  Type="text"  Value="REALA REALB REALC"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Exact stop mode" Range="61" Field="0">
      <ConditionMacro>
        <Variable Name="4015" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4215" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Automatic corner override mode" Range="62" Field="0">
      <ConditionMacro>
        <Variable Name="4015" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4215" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tapping mode" Range="63" Field="0">
      <ConditionMacro>
        <Variable Name="4015" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4215" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Cutting mode" Range="64" Field="0">
      <ConditionMacro>
        <Variable Name="4015" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4215" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Electronic gear box synchronization cancellation" Range="80.8" Field="0">
      <ConditionMacro>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Absolute command" Range="90" Field="0">
      <ConditionMacro>
        <Variable Name="4003" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4203" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="ModeAbsolute" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="Optimizable" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Incremental command" Range="91" Field="0">
      <ConditionMacro>
        <Variable Name="4003" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4203" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="ModeIncremental" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="Optimizable" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Inverse time feed" Range="93" Field="0">
      <ConditionMacro>
        <Variable Name="4005" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4205" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="FeedModeTime" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="Optimizable" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Feed/min feed" Range="94" Field="0">
      <ConditionMacro>
        <Variable Name="4005" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4205" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="FeedModeMinute" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="Optimizable" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Feed/rev feed" Range="95" Field="0">
      <ConditionMacro>
        <Variable Name="4005" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4205" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="FeedModeRevolution" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="Optimizable" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Constant surface speed cancel" Range="97" Field="0">
      <ConditionMacro>
        <Variable Name="4013" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4213" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="RPMMode" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Cylindrical Interpolation" Range="107" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="A" Range="0"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="A"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="B" Range="0"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="B"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="C" Range="0"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="C"/>
        <Macro Name="CylindricalInterpolation" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="Unsupported" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
  </SuperGroup>
  <SuperGroup Name="Cycles">
    <Group Name="G" Type="word" Description="Threading" Range="33" Field="0">
      <ConditionMacro>
        <Variable Name="4001" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4201" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CycleTurnThread" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Drilling cycle, deep" Range="73" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CycleStepUpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="CyclesDeep" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Left handed tapping cycle" Range="74" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesTap" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Fine  boring cycle" Range="76" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesBoreOrient" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CycleBoreShiftFlags" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="0 0 -1 0 1"/>
        </Macro>
        <Macro Name="CycleBoreSpindleOnOffFlags" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="1 1 1 1 0 0 0 0 0 1"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Cycle cancel" Range="80" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="4005" Range="94"/>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="FeedModeMinute" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CyclesCancel" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="4005" Range="95"/>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="FeedModeRevolution" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CyclesCancel" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesCancel" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Drilling cycle" Range="81" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesDrill" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Drilling cycle, face" Range="82" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesFace" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Drilling cycle, deep" Range="83" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_05114[2]" Range="&gt;0"/>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesDeep" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CycleStepMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="CycleStepUpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="CycleStepUpValue" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#MP_05114[2]"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesDeep" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CycleStepMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="CycleStepUpMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="CycleStepUpValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0.04"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tapping cycle" Range="84" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesTap" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tapping cycle, rigid (Series 15 format)" Range="84.2" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesTap" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Tapping cycle, reverse rigid (Series 15 format)" Range="84.3" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesTap" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Boring cycle" Range="85" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesBore" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Boring cycle" Range="86" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesBoreOrient" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CycleBoreSpindleOnOffFlags" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="1 1 1 1 0 0 0 0 0 1"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Back boring cycle" Range="87" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesBoreOrient" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CycleBoreShiftFlags" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="-1 1 0 -1 1"/>
        </Macro>
        <Macro Name="CycleBoreSpindleOnOffFlags" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="0 0 0 1 0 0 0 0 0 1"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Boring cycle, orient spindle" Range="88" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesBoreOrient" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Boring cycle, drag" Range="89" Field="0">
      <ConditionMacro>
        <Variable Name="4009" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4209" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CyclesBoreDrag" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CycleBoreShiftFlags" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="-1 1 -1 0 1"/>
        </Macro>
        <Macro Name="CycleBoreSpindleOnOffFlags" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="0 0 0 1 0 0 0 0 0 1"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Cycle return to initial level" Range="98" Field="0">
      <ConditionMacro>
        <Variable Name="4010" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4210" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="SetCycleRetractLevel" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
          <Override  Type="text"  Value="1"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Cycle return to R level" Range="99" Field="0">
      <ConditionMacro>
        <Variable Name="4010" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4210" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="SetCycleRetractLevel" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
          <Override  Type="text"  Value="2"/>
        </Macro>
      </ConditionMacro>
    </Group>
  </SuperGroup>
  <SuperGroup Name="Variables">
    <Group Name="3000" Type="variable" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="AlarmSignal" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="#3000 Alarm Signal"/>
        </Macro>
        <Macro Name="StopUnconditional" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="3006" Type="variable" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="#3006 Operator Message"/>
        </Macro>
        <Macro Name="StopProgram" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="10001" Type="variable" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="TableIndex" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TableValues" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="SetTableValues" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Tool Length Compensation"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="12001" Type="variable" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="TableIndex" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TableValues" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="SetTableValues" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Cutter Compensation"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="15001" Type="variable" Description="" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="4014" Range="54.2"/>
        <CondVariable Logical="and" Name="4130" Range="1"/>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordXValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="SetAbsWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoord" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordXValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="SetAbsWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoord" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="15002" Type="variable" Description="" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="4014" Range="54.2"/>
        <CondVariable Logical="and" Name="4130" Range="1"/>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordYValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="SetAbsWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoord" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordYValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="SetAbsWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="15003" Type="variable" Description="" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="4014" Range="54.2"/>
        <CondVariable Logical="and" Name="4130" Range="1"/>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordZValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="SetAbsWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoord" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordZValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="SetAbsWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
  </SuperGroup>
  <SuperGroup Name="Registers">
    <Group Name="A" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="A" Range="1"/>
        <CondWord Logical="and" Name="A" Range="2"/>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="7.1"/>
        <Variable Name="107118" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CylindricalDiameter" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$ * 2"  Units="0"/>
          <OverrideWordFormat  InchType="trailing"  InchFormat="3.4" MetricType="trailing" MetricFormat="4.3"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="WorkCoordAValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="50.1"/>
        <Macro Name="MirrorACancel" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51.1"/>
        <Macro Name="MirrorAValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="52"/>
        <Macro Name="WorkCoordAValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="53"/>
        <CondState Logical="and_not" Name="ABS_INC" Value="INCREMENTAL"/>
        <Macro Name="AAxisMachineMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="92"/>
        <Macro Name="WorkCoordAValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="107"/>
        <Variable Name="107118" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CylindricalDiameter" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$ * 2"  Units="0"/>
          <OverrideWordFormat  InchType="trailing"  InchFormat="3.4" MetricType="trailing" MetricFormat="4.3"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="INTERPOLATION" Value="CYLINDRICAL"/>
        <Macro Name="AAxisMotion" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$ * 57.29578 / #107118"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="AAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="B" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="7.1"/>
        <Variable Name="107118" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CylindricalDiameter" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$ * 2"  Units="0"/>
          <OverrideWordFormat  InchType="trailing"  InchFormat="3.4" MetricType="trailing" MetricFormat="4.3"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="WorkCoordBValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="50.1"/>
        <Macro Name="MirrorBCancel" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51.1"/>
        <Macro Name="MirrorBValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="52"/>
        <Macro Name="WorkCoordBValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="53"/>
        <CondState Logical="and_not" Name="ABS_INC" Value="INCREMENTAL"/>
        <Macro Name="BAxisMachineMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="92"/>
        <Macro Name="WorkCoordBValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="107"/>
        <Variable Name="107118" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CylindricalDiameter" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$ * 2"  Units="0"/>
          <OverrideWordFormat  InchType="trailing"  InchFormat="3.4" MetricType="trailing" MetricFormat="4.3"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4102" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4302" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="BAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name=",C" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="CornerMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="CornerValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CornerMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="C" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="7.1"/>
        <Variable Name="107118" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CylindricalDiameter" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$ * 2"  Units="0"/>
          <OverrideWordFormat  InchType="trailing"  InchFormat="3.4" MetricType="trailing" MetricFormat="4.3"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="WorkCoordCValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="50.1"/>
        <Macro Name="MirrorCCancel" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51.1"/>
        <Macro Name="MirrorCValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="52"/>
        <Macro Name="WorkCoordCValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="53"/>
        <CondState Logical="and_not" Name="ABS_INC" Value="INCREMENTAL"/>
        <Macro Name="CAxisMachineMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="92"/>
        <Macro Name="WorkCoordCValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="107"/>
        <Variable Name="107118" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CylindricalDiameter" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$ * 2"  Units="0"/>
          <OverrideWordFormat  InchType="trailing"  InchFormat="3.4" MetricType="trailing" MetricFormat="4.3"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="INTERPOLATION" Value="POLAR"/>
        <Macro Name="PolarYValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="CAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="D" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="43.8"/>
        <Macro Name="Tool3dOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="45-48"/>
        <Macro Name="CutterCompOffsetValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4107" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4307" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="ToolCutterComp" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="SetCutterCompVar" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="{12000+$}"/>
        </Macro>
        <Macro Name="SetCutterCompVar" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="{2400+$}"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="DPRNT" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="FanucDprint" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="E" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="Unsupported" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="F" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Variable Name="4109" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4309" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="FeedRate" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="H" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="45-48"/>
        <Macro Name="ToolLengthOffsetValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="M" Range="98"/>
        <Macro Name="SubStartingSeq" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="M" Range="198"/>
        <Macro Name="SubStartingSeq" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4111" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4311" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="GageOffsetDrivenPoint" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="SetGageOffsetVar" Scan="no" AfterMotion="yes">
          <Override  Type="text"  Value="{2000 + $}"/>
        </Macro>
        <Macro Name="SetGageOffsetVar" Scan="no" AfterMotion="yes">
          <Override  Type="text"  Value="{10000 + $}"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="I" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <Macro Name="CircleCenterX" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <Macro Name="CircleCenterX" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="2 3"/>
        <Macro Name="CircleCenterX" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="4008" Range="43.8"/>
        <Macro Name="PartNormalX" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51"/>
        <Macro Name="RotationPlaneXScale" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$/1000"  Units="0"/>
        </Macro>
        <Variable Name="CGT_G51_SCALE_FACTOR_X" Scan="no" AfterMotion="no">
        </Variable>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68"/>
        <Macro Name="IVector" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and_not" Name="Q"/>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="123"/>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="132"/>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="213"/>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="231"/>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="312"/>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="321"/>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="3"/>
        <CondWord Logical="and" Name="Q" Range="1"/>
        <Macro Name="PlaneAxisVectorX" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="3"/>
        <CondWord Logical="and" Name="Q" Range="2"/>
        <Macro Name="PlaneAxisVectorX" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.4"/>
        <Macro Name="RotationPlaneIncAngle3" Scan="no" AfterMotion="yes">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneIncAngle3" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="3D_TOOL_OFFSET" Value="ON"/>
        <Macro Name="Tool3dXOffset" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="IVector" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="J" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <Macro Name="CircleCenterY" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <Macro Name="CircleCenterY" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="2 3"/>
        <Macro Name="CircleCenterY" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="4008" Range="43.8"/>
        <Macro Name="PartNormalY" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51"/>
        <Macro Name="RotationPlaneYScale" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$/1000"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68"/>
        <Macro Name="JVector" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and_not" Name="Q"/>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Y"/>
        </Macro>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="123"/>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Y"/>
        </Macro>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="132"/>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Y"/>
        </Macro>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="213"/>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Y"/>
        </Macro>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="231"/>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Y"/>
        </Macro>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="312"/>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Y"/>
        </Macro>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="321"/>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Y"/>
        </Macro>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="3"/>
        <CondWord Logical="and" Name="Q" Range="1"/>
        <Macro Name="PlaneAxisVectorY" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="3"/>
        <CondWord Logical="and" Name="Q" Range="2"/>
        <Macro Name="PlaneAxisVectorY" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.4"/>
        <Macro Name="RotationPlaneIncAngle2" Scan="no" AfterMotion="yes">
          <Override  Type="text"  Value="X"/>
        </Macro>
        <Macro Name="RotationPlaneIncAngle2" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="3D_TOOL_OFFSET" Value="ON"/>
        <Macro Name="Tool3dYOffset" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="JVector" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="K" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <Macro Name="CircleCenterZ" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <Macro Name="CircleCenterZ" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="2 3"/>
        <Macro Name="CircleCenterZ" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <CondWord Logical="and" Name="K" Range="0-9999"/>
        <Macro Name="LoopCount" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="4008" Range="43.8"/>
        <Macro Name="PartNormalZ" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="6.2"/>
        <Macro Name="NurbsKnot" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51"/>
        <Macro Name="RotationPlaneYScale" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$/1000"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68"/>
        <Macro Name="KVector" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and_not" Name="Q"/>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="123"/>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="132"/>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="213"/>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="231"/>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="312"/>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and" Name="Q" Range="321"/>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="3"/>
        <CondWord Logical="and" Name="Q" Range="1"/>
        <Macro Name="PlaneAxisVectorZ" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="X"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="3"/>
        <CondWord Logical="and" Name="Q" Range="2"/>
        <Macro Name="PlaneAxisVectorZ" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.4"/>
        <Macro Name="RotationPlaneIncAngle1" Scan="no" AfterMotion="yes">
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="RotationPlaneIncAngle1" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="3D_TOOL_OFFSET" Value="ON"/>
        <Macro Name="Tool3dZOffset" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="NURBS"/>
        <Macro Name="NurbsKnot" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="KVector" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="L" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="4008" Range="43.8"/>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <Macro Name="HelicalFullLoops" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <Macro Name="HelicalFullLoops" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="LoopCount" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="P" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="4"/>
        <Macro Name="DwellTime" Scan="no" AfterMotion="no">
          <OverrideWordFormat  InchType="trailing"  InchFormat="3.4" MetricType="trailing" MetricFormat="4.3"/>
        </Macro>
        <Variable Name="3001" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#3001 + $"  Units="0"/>
        </Variable>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="5"/>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="5.1"/>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="6.2"/>
        <Macro Name="NurbsOrder" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <CondWord Logical="and" Name="P" Range="0"/>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="yes">
          <Override  Type="text"  Value="Base Work Offset 1 1 5200 20"/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="yes">
          <Override  Type="text"  Value="Base Work Offset 1 1 100300 50 "/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <CondWord Logical="and" Name="L" Range="20"/>
        <CondWord Logical="and" Name="P" Range="1-300"/>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="TableIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <CondWord Logical="and" Name="L" Range="21"/>
        <CondWord Logical="and" Name="P" Range="1-8"/>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TableIndex" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <CondWord Logical="and" Name="L" Range="21"/>
        <CondWord Logical="and" Name="P"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="P1 - P8 are valid for G10 L21"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <CondWord Logical="and" Name="L" Range="23"/>
        <CondWord Logical="and" Name="P" Range="1-7"/>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="4"  Units="0"/>
        </Macro>
        <Macro Name="TableIndex" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="yes">
          <Override  Type="text"  Value="Work Offsets 54 59 5220 20"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="30"/>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51"/>
        <Variable Name="CGT_G51_SCALE_FACTOR" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="$/1000"  Units="0"/>
        </Variable>
        <Macro Name="RotationPlaneXScale" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#CGT_G51_SCALE_FACTOR"  Units="0"/>
        </Macro>
        <Macro Name="RotationPlaneYScale" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#CGT_G51_SCALE_FACTOR"  Units="0"/>
        </Macro>
        <Macro Name="RotationPlaneZScale" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#CGT_G51_SCALE_FACTOR"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="54"/>
        <Variable Name="4014" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="54.1"/>
        <Variable Name="4130" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4330" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="54.2"/>
        <CondWord Logical="and" Name="P" Range="0"/>
        <Macro Name="ResetWorkCoordParms" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="SetAbsWorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="SecondaryWorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="ERROR"/>
        </Macro>
        <Macro Name="CancelShiftOffsets" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Variable Name="5500" Scan="no" AfterMotion="no">
        </Variable>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="54.2"/>
        <Variable Name="5500" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="54.4"/>
        <CondWord Logical="and" Name="P" Range="0"/>
        <CondVariable Logical="and" Name="4016" Range="68.2"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="*** G68.2 still active, G69 missing before G54.4 P0 ***"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="54.4"/>
        <CondWord Logical="and" Name="P" Range="0"/>
        <Variable Name="CGT_G54.4_PNO" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsTypes" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="WORK"/>
        </Macro>
        <Macro Name="RpcpOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="SetRelLocationOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WPMoveOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WPLocalOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="WorkingPlaneCancel" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="FIVE"/>
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="54.4"/>
        <CondWord Logical="and" Name="P" Range="1-7"/>
        <CondVariable Logical="and" Name="CGT_G54.4_PNO" Range="&gt;0"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="*** Alarm PS0437 ILLEGAL COMMAND IN WSEC ***"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="54.4"/>
        <CondWord Logical="and" Name="P" Range="1-7"/>
        <Variable Name="CGT_G54.4_PNO" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="DynamicWorkOffsetsTypes" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="WORK"/>
        </Macro>
        <Macro Name="DynamicWorkOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RpcpOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="SetRelLocationOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordIndex" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="4"  Units="0"/>
        </Macro>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#(26003+($*10))"  Units="0"/>
          <Override  Type="text"  Value="X"/>
        </Macro>
        <Macro Name="RotationPlaneAngle2" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#(26004+($*10))"  Units="0"/>
          <Override  Type="text"  Value="Y"/>
        </Macro>
        <Macro Name="RotationPlaneAngle3" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#(26005+($*10))"  Units="0"/>
          <Override  Type="text"  Value="Z"/>
        </Macro>
        <Macro Name="WPMoveOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="WPLocalOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WPCreate" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="REALA REALB REALC "/>
        </Macro>
        <Macro Name="WorkCoordXValue" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#(26000+($*10))"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordYValue" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#(26001+($*10))"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordZValue" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#(26002+($*10))"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordAValue" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#REALA"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordBValue" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#REALB"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoordCValue" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#REALC"  Units="0"/>
        </Macro>
        <Macro Name="SetAbsWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="65 66 66.1"/>
        <Macro Name="SubroutineName" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <Variable Name="CGT_G68.2_P" Scan="no" AfterMotion="no">
        </Variable>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="M" Range="98"/>
        <Macro Name="SubroutineName" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="M" Range="99"/>
        <Variable Name="100016" Scan="no" AfterMotion="no">
        </Variable>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="M" Range="198"/>
        <Macro Name="SubroutineName" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <Macro Name="CycleDwellTime" Scan="no" AfterMotion="no">
          <OverrideWordFormat  InchType="trailing"  InchFormat="3.4" MetricType="trailing" MetricFormat="4.3"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="NURBS"/>
        <Macro Name="NurbsOrder" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="Unsupported" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="PCLOS" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="FanucPclose" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="POPEN" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="FanucPopen" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="DPRNT_report.txt"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="Q" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="2 3"/>
        <Macro Name="RadiusDeltaPerRev" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="5"/>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="5.1"/>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <CondWord Logical="and" Name="P" Range="1"/>
        <CondWord Logical="and_not" Name="Q" Range="123 132 213 231 312 321"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="PS5457 G68.2/G68.3 FORMAT ERROR"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="73 83"/>
        <Macro Name="CycleStepValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="76 87 89"/>
        <Macro Name="CycleBoreShiftValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="M" Range="98"/>
        <Macro Name="SubStartingSeq" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <Macro Name="CycleStepValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="N"/>
        <CondWord Logical="and" Name="Q"/>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="Unsupported" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name=",R" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="CornerMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="CornerValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CornerMode" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="R" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <Macro Name="CircleRadius" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <Macro Name="CircleRadius" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="5"/>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="5.1"/>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="6.2"/>
        <Macro Name="NurbsWeight" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="TableValues" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68"/>
        <Macro Name="RotationAngle" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RotationPlaneAngle1" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <CondState Logical="and" Name="ABS_INC" Value="ABSOLUTE"/>
        <Macro Name="SetCycleRapidLevel" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <CondState Logical="and" Name="ABS_INC" Value="INCREMENTAL"/>
        <Macro Name="SetCycleRapidLevel" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="1"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_TYPE" Value="NURBS"/>
        <Macro Name="NurbsWeight" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="Unsupported" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="S" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="M" Range="19"/>
        <Macro Name="SpindleOrient" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4119" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4319" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="ActiveSpindleSpeed" Scan="no" AfterMotion="no">
          <OverrideWordFormat  InchType="decimal"  InchFormat="5.0" MetricType="decimal" MetricFormat="5.0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="T" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Variable Name="4120" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4320" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="ToolCode" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="U" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="WorkCoordUValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="52"/>
        <Macro Name="WorkCoordUValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="53"/>
        <CondState Logical="and_not" Name="ABS_INC" Value="INCREMENTAL"/>
        <Macro Name="UAxisMachineMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="92"/>
        <Macro Name="WorkCoordUValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="M" Range="6"/>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="UAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="V" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="WorkCoordVValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="52"/>
        <Macro Name="WorkCoordVValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="53"/>
        <CondState Logical="and_not" Name="ABS_INC" Value="INCREMENTAL"/>
        <Macro Name="VAxisMachineMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="92"/>
        <Macro Name="WorkCoordVValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="VAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="W" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="WorkCoordWValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="52"/>
        <Macro Name="WorkCoordWValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="53"/>
        <CondState Logical="and_not" Name="ABS_INC" Value="INCREMENTAL"/>
        <Macro Name="WAxisMachineMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="92"/>
        <Macro Name="WorkCoordWValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="WAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="X" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="0"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MotionLinear" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="XAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="MotionCW" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="0"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MotionLinear" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="XAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="MotionCCW" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="1"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="PS0022, “R OR I,J,K COMMAND NOT FOUND”"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="1"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="PS0022, “R OR I,J,K COMMAND NOT FOUND”"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="4"/>
        <Macro Name="DwellTime" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="WorkCoordXValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="50.1"/>
        <Macro Name="MirrorXCancel" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51"/>
        <Macro Name="RotationPlaneXPoint" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51.1"/>
        <Macro Name="MirrorXValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="52"/>
        <Macro Name="WorkCoordXValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="53"/>
        <CondState Logical="and_not" Name="ABS_INC" Value="INCREMENTAL"/>
        <Macro Name="XAxisMachineMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68"/>
        <Macro Name="RotateXValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RotationPlaneXPoint" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <Macro Name="RotationPlaneXPoint" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoordXValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.4"/>
        <Macro Name="RotationPlaneIncXPoint" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="92"/>
        <Macro Name="WorkCoordXValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <CondState Logical="and" Name="ABS_INC" Value="ABSOLUTE"/>
        <CondState Logical="and" Name="MOTION_PLANE" Value="YZ"/>
        <Macro Name="SetCycleDepth" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="0"/>
        </Macro>
        <Macro Name="SetMotionFlag" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <CondState Logical="and" Name="ABS_INC" Value="INCREMENTAL"/>
        <CondState Logical="and" Name="MOTION_PLANE" Value="YZ"/>
        <Macro Name="SetCycleDepth" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="2"/>
        </Macro>
        <Macro Name="SetMotionFlag" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_PLANE" Value="ZX"/>
        <CondState Logical="and" Name="POLAR_COORD" Value="ON"/>
        <Macro Name="SetPolarAngle" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_PLANE" Value="XY"/>
        <CondState Logical="and" Name="POLAR_COORD" Value="ON"/>
        <Macro Name="SetPolarRadius" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="INTERPOLATION" Value="POLAR"/>
        <Macro Name="PolarXValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="XAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="Y" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="0"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MotionLinear" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="YAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="MotionCW" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="0"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MotionLinear" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="YAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="MotionCCW" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="1"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="PS0022, “R OR I,J,K COMMAND NOT FOUND”"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="1"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="PS0022, “R OR I,J,K COMMAND NOT FOUND”"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="WorkCoordYValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="50.1"/>
        <Macro Name="MirrorYCancel" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51"/>
        <Macro Name="RotationPlaneYPoint" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51.1"/>
        <Macro Name="MirrorYValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="52"/>
        <Macro Name="WorkCoordYValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="53"/>
        <CondState Logical="and_not" Name="ABS_INC" Value="INCREMENTAL"/>
        <Macro Name="YAxisMachineMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68"/>
        <Macro Name="RotateYValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RotationPlaneYPoint" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <Macro Name="RotationPlaneYPoint" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoordYValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.4"/>
        <Macro Name="RotationPlaneIncYPoint" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="92"/>
        <Macro Name="WorkCoordYValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <CondState Logical="and" Name="ABS_INC" Value="ABSOLUTE"/>
        <CondState Logical="and" Name="MOTION_PLANE" Value="ZX"/>
        <Macro Name="SetCycleDepth" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="0"/>
        </Macro>
        <Macro Name="SetMotionFlag" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <CondState Logical="and" Name="ABS_INC" Value="INCREMENTAL"/>
        <CondState Logical="and" Name="MOTION_PLANE" Value="ZX"/>
        <Macro Name="SetCycleDepth" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="2"/>
        </Macro>
        <Macro Name="SetMotionFlag" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_PLANE" Value="XY"/>
        <CondState Logical="and" Name="POLAR_COORD" Value="ON"/>
        <Macro Name="SetPolarAngle" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_PLANE" Value="YZ"/>
        <CondState Logical="and" Name="POLAR_COORD" Value="ON"/>
        <Macro Name="SetPolarRadius" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="YAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="Z" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="0"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MotionLinear" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ZAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="MotionCW" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="0"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MotionLinear" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="ZAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="MotionCCW" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="1"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="PS0022, “R OR I,J,K COMMAND NOT FOUND”"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondVariable Logical="and" Name="MP_03403[5]" Range="1"/>
        <CondState Logical="and" Name="MOTION_TYPE" Value="CCW"/>
        <CondWord Logical="and_not" Name="I"/>
        <CondWord Logical="and_not" Name="J"/>
        <CondWord Logical="and_not" Name="K"/>
        <CondWord Logical="and_not" Name="R"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="PS0022, “R OR I,J,K COMMAND NOT FOUND”"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="10"/>
        <Macro Name="WorkCoordZValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="50.1"/>
        <Macro Name="MirrorZCancel" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51"/>
        <Macro Name="RotationPlaneZPoint" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="51.1"/>
        <Macro Name="MirrorZValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="52"/>
        <Macro Name="WorkCoordZValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="53"/>
        <CondState Logical="and_not" Name="ABS_INC" Value="INCREMENTAL"/>
        <Macro Name="ZAxisMachineMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68"/>
        <Macro Name="RotateZValue" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RotationPlaneZPoint" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.2"/>
        <Macro Name="RotationPlaneZPoint" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoordZValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="68.4"/>
        <Macro Name="RotationPlaneIncZPoint" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="G" Range="92"/>
        <Macro Name="WorkCoordZValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <CondState Logical="and" Name="ABS_INC" Value="ABSOLUTE"/>
        <CondState Logical="and" Name="MOTION_PLANE" Value="XY"/>
        <Macro Name="SetCycleDepth" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="0"/>
        </Macro>
        <Macro Name="SetMotionFlag" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="CYCLE" Value="ON"/>
        <CondState Logical="and" Name="ABS_INC" Value="INCREMENTAL"/>
        <CondState Logical="and" Name="MOTION_PLANE" Value="XY"/>
        <Macro Name="SetCycleDepth" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="2"/>
        </Macro>
        <Macro Name="SetMotionFlag" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_PLANE" Value="YZ"/>
        <CondState Logical="and" Name="POLAR_COORD" Value="ON"/>
        <Macro Name="SetPolarAngle" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="MOTION_PLANE" Value="ZX"/>
        <CondState Logical="and" Name="POLAR_COORD" Value="ON"/>
        <Macro Name="SetPolarRadius" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondState Logical="and" Name="INTERPOLATION" Value="POLAR"/>
        <Macro Name="PolarZValue" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="ZAxisMotion" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
  </SuperGroup>
  <SuperGroup Name="M_Misc">
    <Group Name="M" Type="word" Description="Program stop" Range="0" Field="0">
      <ConditionMacro>
        <Macro Name="StopProgram" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Optional program stop" Range="1" Field="0">
      <ConditionMacro>
        <Macro Name="StopOptional" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Program end" Range="2" Field="0">
      <ConditionMacro>
        <Macro Name="EndProgramRewind" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Tool change" Range="6" Field="0">
      <ConditionMacro>
        <Macro Name="ToolChangeMachineSubroutine" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Coolant mist" Range="7" Field="0">
      <ConditionMacro>
        <Macro Name="CoolantMist" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Coolant flood" Range="8" Field="0">
      <ConditionMacro>
        <Macro Name="CoolantFlood" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Coolant off" Range="9" Field="0">
      <ConditionMacro>
        <Macro Name="CoolantOff" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Rotary axis clamp on/off" Range="10 11" Field="0">
      <ConditionMacro>
        <Macro Name="IgnoreMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Spindle off, orient" Range="19" Field="0">
      <ConditionMacro>
        <Macro Name="ActiveSpindleOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Program end, rewind" Range="30" Field="0">
      <ConditionMacro>
        <Macro Name="ActiveSpindleOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="EndProgramRewindSpecial" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Subroutine call" Range="98" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="P"/>
        <Macro Name="CallSub" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="H"/>
        <CondWord Logical="and_not" Name="P"/>
        <Macro Name="CallSubCurrent" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="&lt;"/>
        <Macro Name="CallSub" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="Unsupported" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Subroutine / NC macro return" Range="99" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="P"/>
        <Macro Name="ReturnFromSub" Scan="no" AfterMotion="yes">
        </Macro>
        <Macro Name="GotoJumpForwardBackwardStart" Scan="no" AfterMotion="yes">
          <Override  Type="expression"  Value="#100016"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="ReturnFromSub" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="Subroutine call " Range="198" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="P"/>
        <Macro Name="CallSub" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="H"/>
        <CondWord Logical="and_not" Name="P"/>
        <Macro Name="CallSubCurrent" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="Unsupported" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="M" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="Unsupported" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
  </SuperGroup>
  <SuperGroup Name="G_Prep">
    <Group Name="G" Type="word" Description="Dwell" Range="4" Field="0">
      <ConditionMacro>
        <Macro Name="DwellMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Data input" Range="10" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="L" Range="20"/>
        <CondVariable Logical="and" Name="4014" Range="54.1"/>
        <CondWord Logical="and" Name="P" Range="1-48"/>
        <Macro Name="SetAdditionalWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#4130"  Units="0"/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 48 2 2 7000 20 20 "/>
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 300 2 2 14000 20 20 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 300 2 2 101000 50 50 "/>
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="L" Range="20"/>
        <CondVariable Logical="and" Name="4014" Range="54.1"/>
        <CondWord Logical="and" Name="P" Range="49-300"/>
        <Macro Name="SetAdditionalWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#4130"  Units="0"/>
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 300 2 2 14000 20 20 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 300 2 2 101000 50 50 "/>
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="L" Range="20"/>
        <CondWord Logical="and" Name="P" Range="1-48"/>
        <Macro Name="SetAdditionalWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 48 2 2 7000 20 20 "/>
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 300 2 2 14000 20 20"/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 300 2 2 101000 50 50 "/>
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="L" Range="20"/>
        <CondWord Logical="and" Name="P" Range="49-300"/>
        <Macro Name="SetAdditionalWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 300 2 2 14000 20 20"/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 300 2 2 101000 50 50 "/>
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="L" Range="21"/>
        <Macro Name="SetAdditionalWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 8 5520 20 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 8 117050 50"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="L" Range="23"/>
        <Macro Name="SetAdditionalWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="4"  Units="0"/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 67 4 4 26009 10 10 1 2 3 4 5 6 -1 -1 -1 "/>
        </Macro>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="L" Range="2"/>
        <CondVariable Logical="and" Name="4014" Range="54-59"/>
        <Macro Name="SetWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="expression"  Value="#4014"  Units="0"/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 54 59 5220 20"/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 54 59 100350 50 "/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="L" Range="2"/>
        <Macro Name="SetWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 54 59 5220 20 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 54 59 100350 50"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="L" Range="10"/>
        <Macro Name="SetTableValues" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Tool Length Compensation"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="L" Range="12"/>
        <Macro Name="SetTableValues" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Cutter Compensation"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Reference point return" Range="28" Field="0">
      <ConditionMacro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="ReferencePointIndex" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Reference point return" Range="30" Field="0">
      <ConditionMacro>
        <CondWord Logical="and_not" Name="P"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="ReferencePointIndex" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="P" Range="2"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="ReferencePointIndex" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="P" Range="3"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="ReferencePointIndex" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="3"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="P" Range="4"/>
        <Variable Name="4008" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="ReferencePointIndex" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="4"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="P"/>
        <Macro Name="ErrorMacro" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="P value out of range for G30 - Defaulting to P2"/>
        </Macro>
        <Variable Name="4008" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Variable Name="4208" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="49"  Units="0"/>
        </Variable>
        <Macro Name="RotaryControlPointOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="ReferencePointIndex" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Cutter/3D tool compensation cancel" Range="40" Field="0">
      <ConditionMacro>
        <Variable Name="4007" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4207" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CutterCompOff" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="Cancel3dToolOffset" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Cutter compensation left/3D tool compensation on" Range="41" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="I"/>
        <Variable Name="4007" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4207" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CutterComp3d" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="J"/>
        <Variable Name="4007" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4207" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CutterComp3d" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="K"/>
        <Variable Name="4007" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4207" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CutterComp3d" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4007" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4207" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CutterCompLeft" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Cutter compensation right" Range="42" Field="0">
      <ConditionMacro>
        <Variable Name="4007" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4207" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CutterCompRight" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Scaling OFF" Range="50" Field="0">
      <ConditionMacro>
        <Variable Name="4011" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4021" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="RotationPlaneCancelReset" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Scaling ON" Range="51" Field="0">
      <ConditionMacro>
        <Variable Name="4011" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4021" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="RotationPlane2" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Local coordinate system" Range="52" Field="0">
      <ConditionMacro>
        <Macro Name="AbsoluteShiftNum" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Additional work offset coordinate system " Range="54" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="5500" Range="&gt; 0"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="G54.2 needs to be cancelled with P0 first"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="P"/>
        <Variable Name="4014" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4214" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="AdditionalWorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="ERROR"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4014" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4214" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="ERROR"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Additional work offset coordinate system" Range="54.1" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="5500" Range="&gt; 0"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="G54.2 needs to be cancelled with P0 first"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4014" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4214" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="CancelShiftOffsets" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="AdditionalWorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="ERROR"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Dynamic work offset coordinate system " Range="54.2" Field="0">
      <ConditionMacro>
        <Variable Name="4023" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4223" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="SecondaryWorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="ERROR"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Workpiece setting error compensation" Range="54.4" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="P" Range="1-7"/>
        <Variable Name="4033" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4233" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="SecondaryWorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="ERROR"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="P" Range="0"/>
        <CondVariable Logical="and" Name="4008" Range="43.4"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
          <Override  Type="text"  Value="G49 must be called before G54.4 P0"/>
        </Macro>
        <Macro Name="StopUnconditional" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="P" Range="0"/>
        <Variable Name="4033" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Variable>
        <Variable Name="4233" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Variable>
        <Macro Name="CancelSecondaryWorkOffset" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Work offset coordinate system 1-6 selection" Range="55-59" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="5500" Range="&gt; 0"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="G54.2 needs to be cancelled with P0 first"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4014" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4214" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="TableSubRegisterValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WorkCoord" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="ERROR"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="NC macro call continuous per block" Range="66.1" Field="0">
      <ConditionMacro>
        <Variable Name="4012" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4212" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="VarsUpdateMacroVars" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="CallNCMacroBlock" Scan="no" AfterMotion="yes">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Coordinate system rotation / 3D Coordinate system rotation" Range="68" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="I"/>
        <Variable Name="4016" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4216" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="FanucRotationPlaneLocal2" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="J"/>
        <Variable Name="4016" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4216" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="FanucRotationPlaneLocal2" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="K"/>
        <Variable Name="4016" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4216" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="FanucRotationPlaneLocal2" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4016" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4216" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="RotateMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Absolute tilted working plane command" Range="68.2" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="P" Range="3"/>
        <CondWord Logical="and" Name="Q" Range="1"/>
        <Variable Name="4016" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4216" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="RotationPlaneInput" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="3"  Units="0"/>
        </Macro>
        <Macro Name="RotationPlane" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RotationPlaneInput" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <CondWord Logical="and" Name="P" Range="3"/>
        <CondWord Logical="and" Name="Q" Range="2"/>
        <Variable Name="4016" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4216" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="RotationPlaneInput" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="3"  Units="0"/>
        </Macro>
        <Macro Name="RotationPlane2" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="UpdateAxisValues" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RotationPlaneInput" Scan="no" AfterMotion="yes">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4016" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4216" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="RotationPlane2" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Incremental tilted working plane command" Range="68.4" Field="0">
      <ConditionMacro>
        <Variable Name="4016" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4216" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="RotationPlane2" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Coordinate system rotation cancel" Range="69" Field="0">
      <ConditionMacro>
        <CondVariable Logical="and" Name="4016" Range="68.2"/>
        <CondVariable Logical="and" Name="4008" Range="43.4"/>
        <Macro Name="MessageMacro" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="G49 must be called before G69"/>
        </Macro>
        <Macro Name="StopUnconditional" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Variable Name="4016" Scan="no" AfterMotion="no">
        </Variable>
        <Variable Name="4216" Scan="no" AfterMotion="no">
        </Variable>
        <Macro Name="RotateCancel" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CancelShiftOffsets" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="WorkingPlaneCancel" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="SIX"/>
        </Macro>
        <Macro Name="RotationPlaneCancelReset" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="FOUR"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="Position set" Range="92" Field="0">
      <ConditionMacro>
        <Macro Name="ChangeWorkCoord" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="G" Type="word" Description="" Field="0">
      <ConditionMacro>
        <Macro Name="Unsupported" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
  </SuperGroup>
  <SuperGroup Name="CGTech">
    <Group Name="OPEN" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="DOORS"/>
        <Macro Name="ChangeSubsystemID" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Door"/>
        </Macro>
        <Macro Name="SaveUnits" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="UnitsMetric" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="UAxisMachineMotion" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RestoreUnits" Scan="no" AfterMotion="yes">
        </Macro>
        <Macro Name="RestoreChangeSubsystemID" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="CLOSE" Type="word" Description="" Field="0">
      <ConditionMacro>
        <CondWord Logical="and" Name="DOORS"/>
        <Macro Name="ChangeSubsystemID" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Door"/>
        </Macro>
        <Macro Name="SaveUnits" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="UnitsMetric" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="UAxisMachineMotion" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="850"  Units="0"/>
        </Macro>
        <Macro Name="ProcessMotion" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="RestoreUnits" Scan="no" AfterMotion="yes">
        </Macro>
        <Macro Name="RestoreChangeSubsystemID" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
      <ConditionMacro>
        <Macro Name="NullMacro" Scan="no" AfterMotion="no">
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="BLANK" Type="word" Description="Set component visibility: blank" Field="0">
      <ConditionMacro>
        <Macro Name="SetComponentVisibility" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
    <Group Name="UNBLANK" Type="word" Description="Set component visibility: both" Field="0">
      <ConditionMacro>
        <Macro Name="SetComponentVisibility" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="3"  Units="0"/>
        </Macro>
      </ConditionMacro>
    </Group>
  </SuperGroup>
  <Events>
    <Event Type="conversion_init">
      <EventSub Name="">
        <Macro Name="AutosetToolLengthCompVars" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2000"  Units="0"/>
        </Macro>
        <Macro Name="AutosetToolLengthCompVars" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="10000"  Units="0"/>
        </Macro>
        <Macro Name="AllowVariableZeroToBeSet" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="BlockSkipAnywhere" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="ConversionInit" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="CutterCompFull" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="GlobalVariables" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="-1"/>
        </Macro>
        <Macro Name="NCVarWarnMessage" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
      </EventSub>
      <EventSub Name="1">
        <Macro Name="AxisVarsMapping" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="{#Mapping}"/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 54 59 5220 20 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Base Work Offset 1 1 5200 20 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Base Work Offset 1 1 100300 50 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 54 59 100350 50 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 8 5520 20 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 48 2 2 7000 20 20"/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 300 2 2 14000 20 20 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 300 2 2 101000 50 50 "/>
        </Macro>
        <Macro Name="AutosetTableAxisVarsAdv" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="Work Offsets 1 67 4 4 26009 10 10 1 2 3 4 5 6 -1 -1 -1 "/>
        </Macro>
        <Macro Name="BiDirVarsAbsIncr" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="BiDirVarsApplyOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="BiDirVarsModeOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="CutterCompOnOffSwitches" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
          <Override  Type="text"  Value="11"/>
        </Macro>
        <Macro Name="CylindricalRotationalAxis" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="4"  Units="0"/>
        </Macro>
        <Macro Name="CylindricalLinearAxis" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="2"  Units="0"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsTypes" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
          <Override  Type="text"  Value="SECONDARY"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsTypes" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
          <Override  Type="text"  Value="SHIFT"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsTypes" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
          <Override  Type="text"  Value="WORK"/>
        </Macro>
        <Macro Name="DynamicWorkOffsetsMode" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="3"  Units="0"/>
        </Macro>
        <Macro Name="GageOffsetAttributes" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="10000010000"/>
        </Macro>
        <Macro Name="G65ResetVarsOnLoop" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="IfCheckAllMacros" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="Ijk2AbcType" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="99"  Units="0"/>
        </Macro>
        <Macro Name="Ijk2AnglesAddOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="Ijk2AnglesInput" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="IJK"/>
        </Macro>
        <Macro Name="Ijk2AnglesLogic" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="Ijk2AnglesPickSolution" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="1 2"/>
        </Macro>
        <Macro Name="Ijk2AnglesVerticalRule" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="SetCyclePartSurface" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="-0.1"  Units="0"/>
          <Override  Type="text"  Value="2"/>
        </Macro>
        <Macro Name="SetCycleRetractLevel" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
          <Override  Type="text"  Value="1"/>
        </Macro>
        <Macro Name="PolarRotationalAxis" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="6"  Units="0"/>
        </Macro>
        <Macro Name="PolarLinearAxis" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="RpcpOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="RotationPlaneSpecialOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="SetCycleSquareOffType" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="SetDynamicVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="CurGageOffset 5081 5082 5083"/>
        </Macro>
        <Macro Name="SetDynamicVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="AxisLocalPlus 0000000 5000"/>
        </Macro>
        <Macro Name="SetDynamicVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="AxisLocalPlus 1000000 5040"/>
        </Macro>
        <Macro Name="SetDynamicVars" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="AxisMachineMinus 0000000 5020"/>
        </Macro>
        <Macro Name="SetRelLocationOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffGagePivotOffset" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="0"  Units="0"/>
        </Macro>
        <Macro Name="TurnOnOffRealTimeClock" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="VariableListArrayAlias" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="5521,_FOFS,8,20"/>
        </Macro>
        <Macro Name="VariableListArrayAlias" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="7001,_WZP,48,20"/>
        </Macro>
        <Macro Name="VariableListArrayAlias" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="101001,_WZP,300,,20"/>
        </Macro>
        <Macro Name="VariableListArrayAlias" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="14001,_WZP,300,20"/>
        </Macro>
        <Macro Name="WPAdjustOffsets" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="WPMoveOnOff" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
        <Macro Name="OptiPathOptimizeCDC" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </EventSub>
    </Event>
    <Event Type="conversion_finish">
      <EventSub Name="">
        <Macro Name="ConversionFinish" Scan="no" AfterMotion="no">
        </Macro>
      </EventSub>
    </Event>
    <Event Type="block_init">
      <EventSub Name="">
        <Macro Name="BlockInit" Scan="no" AfterMotion="no">
        </Macro>
      </EventSub>
    </Event>
    <Event Type="block_finish">
      <EventSub Name="">
        <Macro Name="BlockFinish" Scan="no" AfterMotion="no">
        </Macro>
        <Macro Name="UpdateVariableBlockProcessing" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="3001;#3001 + (#CGT_REAL_TIME * 1000);Milliseconds"/>
        </Macro>
        <Macro Name="UpdateVariableBlockProcessing" Scan="no" AfterMotion="no">
          <Override  Type="text"  Value="3002;#3002 + (#CGT_REAL_TIME / 3600);Hours"/>
        </Macro>
      </EventSub>
    </Event>
    <Event Type="start_init">
      <EventSub Name="">
        <Macro Name="NCVarDefaultValue" Scan="no" AfterMotion="no">
          <Override  Type="value"  Value="1"  Units="0"/>
        </Macro>
      </EventSub>
    </Event>
  </Events>
  <Variables>
    <Variable Name="3001" Type="value" System="" Description="" Expand="off">
      <NumberType>real</NumberType>
      <Vacancy Current="false" Initial="false" />
      <Value>0</Value>
    </Variable>
    <Variable Name="3002" Type="value" System="" Description="" Expand="off">
      <NumberType>real</NumberType>
      <Vacancy Current="false" Initial="false" />
      <Value>0</Value>
    </Variable>
  </Variables>
  <MacroModals>
    <MacroModal Name="mmCONTROLLER_TYPE"  Value="mmCTL_GENERIC"/>
    <MacroModal Name="mmTOLERANCE_VALUE"  Value="0.001"/>
    <MacroModal Name="mmORDER_OF_MATH_OPERATIONS"  Value="mmRULES_OF_PRECEDENCE"/>
    <MacroModal Name="mmDEFAULT_WORD"  Value=""/>
    <MacroModal Name="mmMAX_WIRE_ANGLE"  Value="30"/>
    <MacroModal Name="mmDEFAULT_VOLTAGE"  Value="mmON_TEXT"/>
    <MacroModal Name="mmUV_INCREMENTAL_XY"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmXY_DRIVES_UV"  Value="mmNO_TEXT"/>
    <MacroModal Name="mmWORK_TABLE_XY_OUT_PLANE"  Value="0"/>
    <MacroModal Name="mmDEFAULT_TOOL_COMP_ID"  Value="1"/>
    <MacroModal Name="mmTOOL_NUMBER_METHOD"  Value="mmTOOL_SELECT_ONLY"/>
    <MacroModal Name="mmTOOL_CHANGE_RETRACT_METHOD"  Value="mmTOOL_NO_RETRACT"/>
    <MacroModal Name="mmTOOL_CHANGE_CAUSE_CANCEL"  Value="mmNO_TEXT"/>
    <MacroModal Name="mmTOOL_CHANGE_CAUSE_RAPID"  Value="mmNO_TEXT"/>
    <MacroModal Name="mmDEFAULT_WORK_OFFSET"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmDEFAULT_WORK_INDEX"  Value="54"/>
    <MacroModal Name="mmDEFAULT_MOTION_TYPE"  Value="mmRAPID"/>
    <MacroModal Name="mmDEFAULT_MOTION_PLANE"  Value="mmMOTION_PLANE_XY"/>
    <MacroModal Name="mmDEFAULT_UNITS"  Value="mmUNITS_PROJECT"/>
    <MacroModal Name="mmDEFAULT_ABSINCR"  Value="mmABSOLUTE_MODE"/>
    <MacroModal Name="mmDEFAULT_FEEDMODE"  Value="mmUNITS_PER_MINUTE"/>
    <MacroModal Name="mmDEFAULT_FEEDRATE"  Value="1"/>
    <MacroModal Name="mmOUTPUT_INIT_SPINDLE_LOC"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmRAPID_MOTION_CAUSE_CANCEL"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmLINEAR_MOTION_CAUSE_CANCEL"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmCW_MOTION_CAUSE_CANCEL"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmCCW_MOTION_CAUSE_CANCEL"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmRAPID_FEEDRATE_MOTION"  Value="mmRAPID_TEXT"/>
    <MacroModal Name="mmMAX_CUT_FEED_RATE"  Value="100"/>
    <MacroModal Name="mmOUTPUT_PRECISION"  Value="mmFULL_PRECISION"/>
    <MacroModal Name="mmSPECIFIED_OUTPUT_PRECISION"  Value="3"/>
    <MacroModal Name="mmCIRCLE_CENTER_MODE"  Value="mmINCREMENTAL_MODE"/>
    <MacroModal Name="mmCIRCLE_INC_IJK_METHOD"  Value="mmINC_IJK_FROM_START"/>
    <MacroModal Name="mmPITCH_SPECIFIED_WITH_IJK"  Value="mmNO_TEXT"/>
    <MacroModal Name="mmAPT_CIRCLE_OUTPUT"  Value="mmAPT_CIRCLE"/>
    <MacroModal Name="mmCYCLE_RETRACT_LEVEL"  Value="mmRETRACT_R_POINT_TEXT"/>
    <MacroModal Name="mmCYCLE_CLEARANCE_DISTANCE"  Value="0.1"/>
    <MacroModal Name="mmCYCLE_RAPID_METHOD"  Value="mmCLEARANCE_PLANE"/>
    <MacroModal Name="mmCYCLE_DEPTH_VALUE"  Value="mmG_CODE_DEPENDENT"/>
    <MacroModal Name="mmCYCLE_INCREMENTAL_DEPTH_VALUE"  Value="mmRELATIVE_TO_CYCLE_RAPID_VALUE"/>
    <MacroModal Name="mmCYCLE_RAPID_VALUE"  Value="mmG_CODE_DEPENDENT"/>
    <MacroModal Name="mmCYCLE_INCREMENTAL_RAPID_VALUE"  Value="mmRELATIVE_TO_CYCLE_INITIAL_LEVEL"/>
    <MacroModal Name="mmCYCLE_CANCEL_CAUSE_RAPID"  Value="mmNO_TEXT"/>
    <MacroModal Name="mmCYCLE_IGNORE_ROTARY"  Value="mmNO_TEXT"/>
    <MacroModal Name="mmCYCLE_EXECUTE"  Value="mmON_MOTION"/>
    <MacroModal Name="mmA_AXIS_ROTARY_TYPE"  Value="mmLINEAR"/>
    <MacroModal Name="mmB_AXIS_ROTARY_TYPE"  Value="mmLINEAR"/>
    <MacroModal Name="mmC_AXIS_ROTARY_TYPE"  Value="mmABSOLUTE_360"/>
    <MacroModal Name="mmA2_AXIS_ROTARY_TYPE"  Value="mmLINEAR"/>
    <MacroModal Name="mmB2_AXIS_ROTARY_TYPE"  Value="mmLINEAR"/>
    <MacroModal Name="mmC2_AXIS_ROTARY_TYPE"  Value="mmLINEAR"/>
    <MacroModal Name="mmABSOLUTE_ROTARY_DIRECTION"  Value="mmSHORTEST_DISTANCE"/>
    <MacroModal Name="mmRTCP"  Value="mmNO_TEXT"/>
    <MacroModal Name="mmRTCP_CONTOUR"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmRTCP_WITH_MOTION"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmRTCP_USES"  Value="mmGAGE_PIVOT_OFFSET"/>
    <MacroModal Name="mmCDC_ONOFF_METHOD"  Value="mmRAMP_ONOFF_METHOD"/>
    <MacroModal Name="mmCDC_ONOFF_MOTION"  Value="mmCUTTER_PLANE_MOTION_TEXT"/>
    <MacroModal Name="mmTYPE_OF_SUB_NAMES"  Value="mmNUMERIC_NAMES"/>
    <MacroModal Name="mmAUTO_SEQUENCE_NUMBERING"  Value="mmNO_TEXT"/>
    <MacroModal Name="mmADD_SEQUENCE_NUMBER"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmSEQUENCE_INCREMENT_VALUE"  Value="10"/>
    <MacroModal Name="mmWORD_VALUE_SPACING"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmADD_DATETIME_STAMP"  Value="mmYES_TEXT"/>
    <MacroModal Name="mmOUTPUT_FORMAT"  Value="mmWORD_FORMAT"/>
    <MacroModal Name="mmVNC_TIME_SLICE_VALUE"  Value="100"/>
    <MacroModal Name="mmVNC_OFFSET_X"  Value="0"/>
    <MacroModal Name="mmVNC_OFFSET_Y"  Value="0"/>
    <MacroModal Name="mmVNC_OFFSET_Z"  Value="0"/>
  </MacroModals>
  <SyntaxCheck>
    <VC>10111111110000</VC>
  </SyntaxCheck>
</VcControl>



